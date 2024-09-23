package whatsapp

import (
	"bufio"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/virtouso/WhatsappClientServer/ClientApp/shared"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
	"os"
	"os/signal"
	"syscall"
)

var Client *whatsmeow.Client

func Init() {
	initializeDatabase()
	initializeClient()

	go getUserInput()
	handleShutdownSignal()
	Client.Disconnect()
}

func initializeClient() {
	if Client.Store.ID == nil {
		startLoginProcess()
	} else {
		err := Client.Connect()
		if err != nil {
			panic(err)
		}
	}
}

func startLoginProcess() {
	qrChan, _ := Client.GetQRChannel(context.Background())
	err := Client.Connect()
	if err != nil {
		panic(err)
	}

	for evt := range qrChan {
		handleLoginEvent(evt)
	}
}

func handleShutdownSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func handleLoginEvent(evt whatsmeow.QRChannelItem) {
	if evt.Event == "code" {
		fmt.Println("QR code:", evt.Code)
		shared.RenderQRCodeInTerminal(evt.Code)
	} else {
		fmt.Println("Login event:", evt.Event)
	}
}

// log_db_type = postgress
// log_db_con= user=postgres password=moeen777 dbname=whatsapp sslmode=disable
//log_db_con= postgres://postgres:moeen777@localhost/whatsapp?sslmode=disable
//user_db_con= postgres://postgres:moeen777@localhost/users?sslmode=disable

func initializeDatabase() {
	dbLog := waLog.Stdout("Database", "ERROR", true)

	container, err := sqlstore.New(os.Getenv(shared.LogDbTypeKey), os.Getenv(shared.LogDbConKey), dbLog)
	if err != nil {
		panic(err)
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	initializeClientWithDevice(deviceStore)
}

func initializeClientWithDevice(deviceStore *store.Device) {
	clientLog := waLog.Stdout("Client", "ERROR", true)
	Client = whatsmeow.NewClient(deviceStore, clientLog)
	Client.AddEventHandler(receiveMessageEventHandler)
}

func SendMessage(recipient string, messageText string) error {
	jid, err := types.ParseJID(recipient)
	if err != nil {
		return fmt.Errorf("invalid JID %s: %v", recipient, err)
	}

	msg := &waProto.Message{
		Conversation: proto.String(messageText),
	}

	_, err = Client.SendMessage(context.Background(), jid, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	fmt.Println("Message sent to", recipient)
	return nil
}

func receiveMessageEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

func getUserInput() {
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter receiver JID (e.g., 905431070120@s.whatsapp.net): ")
		receiver, _ := reader.ReadString('\n')
		receiver = receiver[:len(receiver)-1] // Trim newline
		receiver = receiver + "@s.whatsapp.net"
		fmt.Print("Enter message to send: ")
		message, _ := reader.ReadString('\n')
		message = message[:len(message)-1] // Trim newline
		SendMessage(receiver, message)
	}
}
