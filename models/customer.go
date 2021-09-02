package models

import (
	"fmt"
	"log"
	"math/rand"

	"mobile-rest-api/config"
	"mobile-rest-api/db"
	"mobile-rest-api/libs/validator"
	u "mobile-rest-api/utils"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	serviceID    = 1
	lenPaymentID = 6
)

//Customer model
type Customer struct {
	//gorm.Model
	ID            uint64    `gorm:"AUTO_INCREMENT" json:"id"`
	PhoneNo       string    `gorm:"size:15; unique_index; not null" json:"phone_no"`
	Name          string    `gorm:" default:'Без имени '" json:"name"`
	Image         string    `json:"image"`
	SmsCode       string    `gorm:"not null" json:"sms_code"`
	Balance       float64   `sql:"type:decimal(12,2)" gorm:"not null; default:'0.00'" json:"balance"`
	Status        int       `gorm:"not null; default:'0'" json:"status"` // 0-non-active, 1-active, 2 - waiting activation, 3 - waiting deactivation
	ServiceID     int       `gorm:"not null; default:'1'"  json:"service_id"`
	DateCreate    time.Time `json:"date_create"`
	Password      string    `gorm:"" json:"password"`
	Surname       string    `json:"surname"`
	BirthDate     time.Time `json:"birth_date"`
	Email         string    `json:"email"`
	Gender        string    `json:"gender"`
	IsIdentified  string    `json:"is_identified"`
	ImagePassport string    `json:"image_passport"`
	Token         string    `json:"token" sql:"-"`
	PaymentID     string    `json:"payment_id"`
}


//Token JWT
type Token struct {
	UserID  uint64
	PhoneNo string
	SmsCode string
	jwt.StandardClaims
}



//CreateCustomer to db
func CreateCustomer(phoneNo string) (resp map[string]interface{}, cust Customer) { //Create User

	var rs = validator.ValidatePhone(phoneNo)
	if !rs {
		resp := u.Message(709, "Введите номер телефона, пример: 900112233")
		return resp, cust
	}
	var smsCode = randomGeneratePassword(4)

	cust = CreateOrUpdate(phoneNo, smsCode)

	resp = u.Message(0, "Аккаунт успешно создан ждите смс код.")
	resp["phone_no"] = cust.PhoneNo[4:]
	return resp, cust

}


func checkPaymentID(paymentID string) bool {
	db.Init()
	var db = db.GetDB()
	defer db.Close()
	var cust Customer
	if err := db.Where("payment_id = ? ", paymentID).First(&cust).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return true
		}
		return false
	}
	return false
}

//CreateOrUpdate to db
func CreateOrUpdate(phoneNo, smsCode string) Customer {
	db.Init()
	var db = db.GetDB()
	defer db.Close()
	var cust Customer
	var res Customer

	if err := db.Where("phone_no = ? AND service_id = ?", phoneNo, serviceID).First(&cust).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			var customer Customer
			customer.PhoneNo = phoneNo
			customer.DateCreate = time.Now()
			customer.Status = 2
			customer.ServiceID = serviceID
			customer.SmsCode = smsCode
			customer.PaymentID = generatePaymentID(lenPaymentID)

			db.Table("customers").Create(&customer)
			res = customer
		}
	} else {
		//log.Println(cust)
		cust.PhoneNo = phoneNo
		cust.DateCreate = time.Now()
		cust.SmsCode = smsCode
		////////////////////////////////////
		if len(cust.PaymentID) != lenPaymentID {
			PaymentID := generatePaymentID(lenPaymentID)
			cust.PaymentID = PaymentID
		}
		///////////////////
		db.Table("customers").Save(&cust)
		res = cust
	}

	log.Println(res)

	return res
}

func generatePaymentID(n int) string {
	var PaymentID string
forgotog:
	PaymentID = randomGeneratePassword(n)
	err := checkPaymentID(PaymentID)
	if !err {
		goto forgotog
	}
	return PaymentID
}

//CheckSmsCodeCustomer func
func CheckSmsCodeCustomer(phone, smsCode string) map[string]interface{} {

	db.Init()
	db := db.GetDB()
	defer db.Close()
	var customer Customer
	var err = db.Table("customers").Where("phone_no = ? AND  sms_code = ? AND service_id = ?", phone, smsCode, serviceID).First(&customer).Error
	if err != nil {
		return u.Message(708, "СМС код неверный !!!")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(smsCode), bcrypt.DefaultCost)

	log.Println("customer>>>>>>>>>>", customer)

	/* err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(sms_code))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Неверные введен номер телефона или смс код. Пожалуйста, попробуйте еще раз")
	} */
	customer.Status = 1
	customer.Password = string(hashedPassword)
	tk := &Token{UserID: customer.ID, PhoneNo: customer.PhoneNo, SmsCode: customer.SmsCode}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	config.Init()
	var cfg = config.Get()
	tokenString, _ := token.SignedString([]byte(cfg.SecretKey))
	customer.Token = tokenString

	db.Table("customers").Save(&customer)
	var resp = u.Message(0, "Успешно !!!")
	customer.Password = ""
	customer.PhoneNo = customer.PhoneNo[4:]
	resp["customer"] = customer
	return resp
}

//CheckAuthCustomer func
func CheckAuthCustomer(customer interface{}) bool {
	if customer == nil {
		return false
	}
	var phoneNo, ID, smsCode, err = getValueFromContext(customer)
	if err {
		return false
	}
	if !CheckCustomerFromToken(ID, phoneNo, smsCode) {
		return false
	}
	return true
}

func getValueFromContext(ctx interface{}) (phoneNo, ID, smsCode string, err bool) {
	var s = (fmt.Sprintf("%v", ctx))
	s = s[1 : len(s)-1]
	ID = strings.Split(s, " ")[0]
	phoneNo = strings.Split(s, " ")[1]
	smsCode = strings.Split(s, " ")[2]
	if ID == "" || phoneNo == "" {
		err = true
		return
	}
	return
}

//CheckCustomerFromToken f
func CheckCustomerFromToken(ID, phoneNo, smsCode string) bool {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var customer Customer
	var err = db.Table("customers").Where("phone_no = ? AND  id = ? AND service_id = ?", phoneNo, ID, 1).First(&customer).Error
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(smsCode))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}

func randomGeneratePassword(n int) string {
	var str string
generate:
	rand.Seed(time.Now().UnixNano())
	str += strconv.Itoa(rand.Intn(9999999999))
	if len(str) < n {
		goto generate
	}
	if len(str) > n {
		str = str[:n]
	}
	return str
}

//GetCustomer f
func GetCustomer(phoneNo string) Customer {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var customer Customer
	err := db.Table("customers").Where("phone_no = ? AND service_id = ?", phoneNo, serviceID).First(&customer).Error
	if err != nil {
		customer.PhoneNo = "NO DATA FOUND"
		return customer
	}

	return customer

}

//Update data
func (customer Customer) Update() Customer {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Save(&customer)
	return customer
}
