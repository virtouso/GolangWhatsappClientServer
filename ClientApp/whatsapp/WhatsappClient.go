package whatsapp

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/virtouso/WhatsappClientServer/ClientApp/shared"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
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

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

func Init() {

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("postgres", "user=postgres password=moeen777 dbname=whatsapp sslmode=disable", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	Client = whatsmeow.NewClient(deviceStore, clientLog)
	Client.AddEventHandler(eventHandler)

	if Client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := Client.GetQRChannel(context.Background())
		err = Client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {

				fmt.Println("QR code:", evt.Code)
				shared.RenderQRCodeInTerminal(evt.Code)

			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = Client.Connect()
		if err != nil {
			panic(err)
		}

	}
	SendMessage(Client, "905431070120@s.whatsapp.net", "hi dfsd fsd fsd fds ")
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	Client.Disconnect()
}

func SendMessage(client *whatsmeow.Client, recipient string, messageText string) error {
	jid, err := types.ParseJID(recipient)
	if err != nil {
		return fmt.Errorf("invalid JID %s: %v", recipient, err)
	}

	msg := &waProto.Message{
		Conversation: proto.String(messageText),
	}

	_, err = client.SendMessage(context.Background(), jid, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	fmt.Println("Message sent to", recipient)
	return nil
}
