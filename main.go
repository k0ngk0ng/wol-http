package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	wol "github.com/sabhiram/go-wol"
)

func wake(macAddr string, bcastAddr string) error {
	if len(macAddr) == 0 {
		return errors.New("<macAddr> is empty")
	}

	if len(bcastAddr) == 0 {
		bcastAddr = "255.255.255.255"
	}

	var localAddr *net.UDPAddr

	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr+":9")
	if err != nil {
		return err
	}

	// Build the magic packet.
	mp, err := wol.New(macAddr)
	if err != nil {
		return err
	}

	// Grab a stream of bytes to send.
	bs, err := mp.Marshal()
	if err != nil {
		return err
	}

	// Grab a UDP connection to send our packet of bytes.
	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("Attempting to send a magic packet to MAC %s\n", macAddr)
	fmt.Printf("... Broadcasting to: %s\n", bcastAddr)
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
	return nil
}

func main() {
	r := gin.Default()
	r.GET("/wol/wake", func(c *gin.Context) {
		mac := c.Query("mac")
		bcast := c.Query("bcast")

		err := wake(mac, bcast)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run(":8000")
}
