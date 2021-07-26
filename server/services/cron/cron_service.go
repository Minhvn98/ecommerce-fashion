package cron

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	ctrl "github.com/Minhvn98/ecommerce-fashion/controllers"
	"github.com/Minhvn98/ecommerce-fashion/models"
	"github.com/Minhvn98/ecommerce-fashion/services/email"
	"github.com/robfig/cron/v3"
)

var C *cron.Cron

func init() {
	C = cron.New(cron.WithSeconds())
}

func SendEmailAboutRevenueOfTheDay() {
	C.AddFunc("00 00 22 * * *", func() {
		date := strings.Split(time.Now().String(), " ")[0]
		subject := date + " - Thông báo về đơn hàng hôm nay"
		body := ""
		order := ctrl.GetNumberOrderToday()
		fmt.Println(order)
		if len(order) == 0 {
			body += "Hôm nay không có đơn hàng nào"
		} else {
			for key, value := range order {
				body += key + ":" + strconv.Itoa(value) + ", "
			}
		}

		email.SendMail("hme57646@eoopy.com", models.Mail{Subject: subject, Body: body})
		fmt.Println("Email sent to")
	})
}
