package belajar

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/Befous/BackendGin/models"
)

func ReturnStruct(DataStuct any) (result string) {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func ReturnString(geojson []FullGeoJson) string {
	var names []string
	for _, geojson := range geojson {
		names = append(names, geojson.Properties.Name)
	}
	result := strings.Join(names, ", ")
	return result
}

// ----------------------------------------------------------------------- User -----------------------------------------------------------------------

func Authorization(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	var auth User
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenname := DecodeGetName(os.Getenv(publickey), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)

	auth.Username = tokenusername

	if tokenname == "" || tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mconn, collname, auth) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	response.Message = "Berhasil decode token"
	response.Status = true
	response.Data.Name = tokenname
	response.Data.Username = tokenusername
	response.Data.Role = tokenrole

	return ReturnStruct(response)
}

func Registrasi(publickey, mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickey), header)
	tokenrole := DecodeGetRole(os.Getenv(publickey), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mconn, collname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "owner" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	if UsernameExists(mconn, collname, datauser) {
		response.Message = "Username telah dipakai"
		return ReturnStruct(response)
	}

	hash, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		response.Message = "Gagal hash password: " + hashErr.Error()
		return ReturnStruct(response)
	}

	datauser.Password = hash

	InsertUser(mconn, collname, datauser)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func Login(privatekey, mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	if !UsernameExists(mconn, collname, datauser) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if !IsPasswordValid(mconn, collname, datauser) {
		response.Message = "Password Salah"
		return ReturnStruct(response)
	}

	user := FindUser(mconn, collname, datauser)

	tokenstring, tokenerr := Encode(user.Name, user.Username, user.Role, os.Getenv(privatekey))
	if tokenerr != nil {
		response.Message = "Gagal encode token: " + tokenerr.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = "Berhasil login"
	response.Token = tokenstring

	return ReturnStruct(response)
}

// VerifyOTPEndpoint handles OTP verification
func VerifyOTPEndpoint(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		OTP      string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Invalid input"})
		return
	}

	// Ambil secret dari database atau tempat penyimpanan lainnya
	// Misalnya: secret := GetUserSecret(request.Username)

	// Contoh, ganti dengan logika pengambilan secret yang benar
	secret := "your-user-secret-key" // Ganti dengan secret yang diambil dari DB

	if VerifyOTP(secret, request.OTP) {
		c.JSON(http.StatusOK, Response{Status: true, Message: "OTP verified successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, Response{Status: false, Message: "Invalid OTP"})
	}
}

//----------------------------------------------------------------------- User OTP-----------------------------------------------------------------------

func GenerateQR(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, Response{Message: "Username is required"})
		return
	}

	secret, qrCodeURL, err := GenerateOTPSecret(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Message: "Error generating OTP secret"})
		return
	}

	c.JSON(http.StatusOK, QRCodeResponse{
		Status:  true,
		Message: "QR code generated",
		OTP:     qrCodeURL,
		Secret:  secret,
	})
}