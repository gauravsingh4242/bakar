package controller

import (
	"github.com/gauravsingh4242/bakar/logger"
	"github.com/gauravsingh4242/bakar/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var userToConnectionMap map[string]*websocket.Conn

func init() {
	userToConnectionMap = make(map[string]*websocket.Conn)
}

func releaseConnection(conn *websocket.Conn, userId string) {
	conn.Close()
	delete(userToConnectionMap, userId)
}

func storeConnectionMapping(conn *websocket.Conn, userId string) {
	userToConnectionMap[userId] = conn
}

func (c *BakarController) WebSocketController(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	if userId == "" {
		logger.Log.Errorf("userId is empty")
		ctx.JSON(400, "userId not present in request header")
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Log.Error("Error upgrading connection to websocket", err.Error())
		ctx.JSON(500, err.Error())
		return
	}
	storeConnectionMapping(conn, userId)
	logger.Log.Infof("WS connection initialised for user: %s", userId)
	defer releaseConnection(conn, userId)

	var inboundMessage *models.InboundMessage

	for {
		err = conn.ReadJSON(&inboundMessage)
		if err != nil {
			logger.Log.Error("not able to read inboundMessage, terminating connection", err.Error())
			return
		}
		if inboundMessage.TerminateConnection {
			logger.Log.Info("client requested to terminate connection. Terminating...")
			break
		}
		logger.Log.Infof("inboundMessage: %s, receiver: %s", inboundMessage.Message, inboundMessage.ReceiverId)
		receiverConnection, ok := userToConnectionMap[inboundMessage.ReceiverId]
		if !ok {
			logger.Log.Infof("user: %s is not online. Not sending inboundMessage", inboundMessage.ReceiverId)
			continue
		}

		err = receiverConnection.WriteJSON(&models.OutboundMessage{
			SenderId: userId,
			Message:  inboundMessage.Message,
		})
		if err != nil {
			logger.Log.Errorf("%s mnable to send message to %s. Error: %s", userId, inboundMessage.ReceiverId, inboundMessage.Message)
			continue
		}
	}

	return
}
