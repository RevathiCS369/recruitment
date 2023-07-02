package start

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"recruit/ent"
	"recruit/ent/employeemaster"
	"recruit/ent/usermaster"
	"strconv"
	"strings"
	"time"
	//"github.com/twilio/twilio-go"
	//"github.com/twilio/twilio-go/rest/api/v2010"
	//"github.com/twilio/twilio-go/v3"
)

type UserExistence struct {
	NewUser bool `json:"newuser"`
}

// "strings"
func QueryUserMasterByUserName(ctx context.Context, client *ent.Client, username string) (string, error) {
	//Can use GetX as well
	if username == "" {
		return "", errors.New("The username is empty...")
	}
	exists, err := client.UserMaster.Query().
		Where(usermaster.UserNameEQ(username)).
		Exist(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to check username existence: %v", err)
	}
	if exists {
		return "Existing User", nil
	}
	return "", fmt.Errorf("No such username: %s", username)
}

func QueryUserMasterByEmpId(ctx context.Context, client *ent.Client, empid int64) (bool, error) {
	if empid == 0 {
		return false, errors.New("Please pass Emp ID as a non-zero parameter")
	}
	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check empid existence: %v", err)
	}
	return !exists, nil
}

func CreateUserByEmpId(ctx context.Context, client *ent.Client, empid int64) (*ent.UserMaster, error) {
	if empid <= 0 {
		return nil, errors.New("Please input a nonzero Emp ID as an input parameter")
	}

	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check empid existence: %v", err)
	}
	if exists {
		user, err := client.UserMaster.Query().
			Where(usermaster.EmployeeIDEQ(empid)).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query existing user: %v", err)
		}

		if user.Password != "" && user.Status {
			user.NewPasswordRequest = false
			user, err = user.Update().SetNewPasswordRequest(false).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update existing user: %v", err)
			}
			return user, nil
		}

		return user, nil
	}
	employee, err := client.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(empid)).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("The Employee Id is not available in Employee Master: %v", err)
	}

	if employee.FacilityID == "" {
		return nil, errors.New("Please update your Office ID with your parent office")
	}

	if employee.MobileNumber == "" {
		return nil, errors.New("Please get your mobile number updated with your parent Office")
	}
	if employee.EmailID == "" {
		return nil, errors.New("Please get your Email ID updated with your parent Office")
	}

	// Trigger the SMS OTP.
	if employee.FacilityID != "" && employee.MobileNumber != "" && employee.EmailID != "" && employee.EmployeeName != "" {
		user, err := client.UserMaster.
			Create().
			SetEmployeeID(empid).
			SetRoleUserCode(1).
			SetUserName(strconv.Itoa(int(empid))).
			SetStatus(false).
			SetMobile(employee.MobileNumber).
			SetEmailID(employee.EmailID).
			SetEmployeeName(employee.EmployeeName).
			SetFacilityID(employee.FacilityID).
			SetNewPasswordRequest(true).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}

		return user, nil
	}

	return nil, nil
}

func QueryUsersByEmpId(ctx context.Context, client *ent.Client, empid int64) (*ent.UserMaster, error) {
	//Can use GetX as well

	user, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).
		Only(ctx)

	if err != nil {
		log.Println("error at gettting users with emp id: ", err)
		return nil, fmt.Errorf("failed at users in user master: %w", err)
	}
	log.Println("user returned by empid : ", user)
	return user, nil
}

// Generate OTP
func GenerateOTP() int32 {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	otp := rand.Intn(max-min+1) + min
	return int32(otp)
}

// Send SMS
func sendSMSs(msg, phone, templateID, entityID, appName string) error {
	url := "https://api.cept.gov.in/sendsms/api/values/sendsms"

	payload := `{
		"Msg": "` + msg + `",
		"Phone": "` + phone + `",
		"TemplateID": "` + templateID + `",
		"EntityID": "` + entityID + `",
		"AppName": "` + appName + `"
	}`

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	log.Println("Response body:", string(body))

	fmt.Println("Response body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	response := string(body)
	if !strings.Contains(response, "SMS Pushed to NIC Successfully") {
		return fmt.Errorf("failed to send SMS")
	}

	return nil
}

// Update Password
func UpdateUserByEmpID(client *ent.Client, empID int64) (*ent.UserMaster, error) {
	ctx := context.Background()
	if empID <= 0 {
		return nil, errors.New("Please input a nonzero Emp ID as an input parameter")
	}

	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		//Where(usermaster.PasswordIsNil()).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check empID existence: %v", err)
	}
	if exists {
		userMaster, err := client.UserMaster.Query().
			Where(usermaster.EmployeeIDEQ(empID)).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve userMaster: %v", err)
		}

		otp := GenerateOTP() // Assuming you have a function to generate the OTP
		//password := userMaster.Password // Assuming you have a function to generate the password

		//err = SendOTPViaSMS(otp, userMaster.Mobile)
		if err != nil {
			return nil, fmt.Errorf("failed to send OTP via SMS: %v", err)
		}
		updatedUser, err := userMaster.Update().
			SetOTP(otp).
			SetPassword(userMaster.Password).
			SetStatus(true).
			SetNewPasswordRequest(false).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %v", err)
		}

		return updatedUser, nil
	}

	return nil, errors.New("The Employee ID is unavailable ")
}

// Validate Login

func ValidateLoginUser(client *ent.Client, newuser *ent.UserMaster) (*ent.UserMaster, error) {
	// Check if the username exists
	ctx := context.Background()

	username := newuser.UserName
	//password := newuser.Password

	if newuser.UserName == "" {
		return nil, errors.New("Username cannot be empty")
	}
	if newuser.Password == "" {
		return nil, errors.New("Password cannot be empty")
	}

	// Check if the username exists
	exists, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(username)).
		Exist(ctx)

	if !exists {
		return nil, fmt.Errorf("no such username: %s", username)
	}

	// Retrieve the user record with the provided username and password
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserName(username)).
		Only(ctx)
	if user.Password != newuser.Password {
		return nil, errors.New("Incorrect Password")
	}
	if err != nil {
		return nil, fmt.Errorf("Incorrect Password")
	}

	return user, nil
}

//sennd sms

//Update Password for an user with EmpID. in main

func UpdateUserPasswordByEmpID(client *ent.Client, empID int64, newUser *ent.UserMaster) (*ent.UserMaster, error) {
	ctx := context.Background()
	if empID <= 0 {
		return nil, errors.New("Please input a nonzero Emp ID as an input parameter")
	}

	userMaster, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("The Employee ID does not exist")
		}
		return nil, fmt.Errorf("failed to retrieve userMaster: %v", err)
	}

	// Check if the newuser.OTP matches the OTP in the database
	if newUser.OTP != userMaster.OTP {
		//log.Printf("Mismatched OTP for Employee ID %d. Expected OTP: %s, Provided OTP: %s", empID, userMaster.OTP, newUser.OTP)
		return nil, errors.New("Invalid OTP")
	}

	/*// Check if OTP has expired (OTP expires + 6 minutes <= current time)
	expirationTime := userMaster.OTPExpires.Add(6 * time.Minute)
	if time.Now().After(expirationTime) {
		return nil, errors.New("OTP has expired")
	}*/

	// Check if the newuser.OTP = payload.OTP, OTP expires+6mins <=currentime and time. (OTP and OTP expires time i)
	if userMaster.Password == "" {
		userMaster, err = userMaster.Update().
			SetPassword(newUser.Password).
			SetOTP(newUser.OTP).
			SetNewPasswordRequest(false).
			SetStatus(true).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %v", err)
		}
		return userMaster, nil

	}

	return nil, errors.New("The Employee has a valid password. Can opt for password reset")
}

func getUserMobileNumber(empID int64) string {
	ctx := context.Background()
	client := ent.NewClient() // Initialize your ent client here

	// Query the UserMaster entity based on the employee ID
	userMaster, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		Only(ctx)
	if err != nil {
		// Handle the error
		return ""
	}

	// Retrieve the mobile number from the UserMaster entity
	mobileNumber := userMaster.Mobile

	return mobileNumber
}

// Admin Login
// Validate username  pswd against db , then send sms if matches. Return PasswordMatch status : true/false.
// Validate OTP and return Login status  : true/ false.
func ValidateAdminUserLogin(client *ent.Client, newUser *ent.UserMaster) (*ent.UserMaster, error) {
	// Check if the newUser and password are not nil
	if newUser == nil {
		return nil, errors.New("UserMaster cannot be nil")
	}

	// Trim the username and password
	newUser.UserName = strings.TrimSpace(newUser.UserName)
	newUser.Password = strings.TrimSpace(newUser.Password)

	// Check if the username exists
	ctx := context.Background()
	exists, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName)).
		Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to check username existence: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("Invalid Username or username not found: %s", newUser.UserName)
	}

	// Retrieve the user record with the provided username
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Compare the password from the input with the user's password stored in the database
	if user.Password != newUser.Password {
		return nil, errors.New("Incorrect Password")
	}

	/// Send SMS and save OTP
	otp, _, err := SendSMSAndSaveOTP(ctx, client, user)
	if err != nil {
		return nil, fmt.Errorf("failed to send SMS and save OTP: %v", err)
	}

	log.Printf("Generated OTP: %s, Mobile Number: %s, Username: %s", otp, user.Mobile, user.UserName)

	return user, nil
}

// validate login  status after pswd matcches , with otp.

func ValidateAdminLogin(client *ent.Client, newUser *ent.UserMaster) (*ent.UserMaster, error) {
	// Check if the newUser and password are not nil
	if newUser == nil {
		return nil, errors.New("UserMaster cannot be nil")
	}

	// Trim the username and password
	newUser.UserName = strings.TrimSpace(newUser.UserName)

	// Check if the username exists
	ctx := context.Background()
	exists, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName)).
		Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to check username existence: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("Invalid Username or username not found: %s", newUser.UserName)
	}

	// Retrieve the user record with the provided username
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Retrieve the OTP saved in the database
	dbOTP := user.OTP

	// Compare the input OTP with the OTP from the database
	if newUser.OTP != dbOTP {
		// Log the mismatched OTPs
		log.Printf("Input OTP: %d, Database OTP: %d, Mobile Number: %s, Username: %s", newUser.OTP, dbOTP, user.Mobile, user.UserName)
		return nil, errors.New("Incorrect OTP")
	}

	return user, nil
}

// User creation & OTP sending

func NewCreateUserByEmpId(ctx context.Context, client *ent.Client, newUsers *ent.UserMaster) (*ent.UserMaster, error) {
	if newUsers.EmployeeID <= 0 {
		return nil, errors.New("Please input a nonzero Emp ID as an input parameter")
	}

	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(newUsers.EmployeeID)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check empid existence: %v", err)
	}
	if exists {
		user, err := client.UserMaster.Query().
			Where(
				usermaster.EmployeeIDEQ(newUsers.EmployeeID),
				//usermaster.PasswordNEQ(""),
				//usermaster.StatusEQ(true),
			).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query existing user: %v", err)
		}

		if user != nil {
			return nil, errors.New("The user already exists")
		}
	}

	employee, err := client.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(newUsers.EmployeeID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("The Employee Id is not available in Employee Master: %v", err)
	}

	if employee.FacilityID == "" {
		return nil, errors.New("Please update your Office ID with your parent office")
	}

	if employee.MobileNumber == "" {
		return nil, errors.New("Please get your mobile number updated with your parent Office")
	}
	if employee.EmailID == "" {
		return nil, errors.New("Please get your Email ID updated with your parent Office")
	}

	// Trigger the SMS OTP.
	if employee.FacilityID != "" && employee.MobileNumber != "" && employee.EmailID != "" && employee.EmployeeName != "" {
		user, err := client.UserMaster.
			Create().
			SetEmployeeID(newUsers.EmployeeID).
			SetRoleUserCode(1).
			SetUserName(strconv.Itoa(int(newUsers.EmployeeID))).
			SetStatus(false).
			SetMobile(employee.MobileNumber).
			SetEmailID(employee.EmailID).
			SetEmployeeName(employee.EmployeeName).
			SetFacilityID(employee.FacilityID).
			//SetNewPasswordRequest(true).
			//SetPassword(newUsers.Password).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}

		// Call trigger SMS function
		otp, _, err := SendSMSAndSaveOTP(ctx, client, user)
		if err != nil {
			return nil, fmt.Errorf("failed to send SMS and save OTP: %v", err)
		}

		log.Printf("Generated OTP: %s, Mobile Number: %s, Username: %s", otp, user.Mobile, user.UserName)

		return user, nil
	}

	return nil, nil
}

// Update First Time username password
func NewValidateLoginUser(client *ent.Client, newuser *ent.UserMaster) (*ent.UserMaster, error) {
	ctx := context.Background()

	// Check if the username exists
	username := newuser.UserName
	if username == "" {
		return nil, errors.New("Username cannot be empty")
	}
	if newuser.Password == "" {
		return nil, errors.New("Password cannot be empty")
	}

	// Check if the username exists
	exists, err := client.UserMaster.
		Query().
		Where(
			usermaster.UserName(username)).
		//usermaster.PasswordNEQ("")).
		Exist(ctx)

	if !exists {
		return nil, fmt.Errorf("no such username: %s", username)
	}

	// Retrieve the user record with the provided username
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserName(username)).
		Only(ctx)
	if err != nil {
		return nil, errors.New("Failed to retrieve user")
	}

	// Validate the input OTP with the stored OTP
	if user.OTP != newuser.OTP {
		return nil, errors.New("Invalid OTP")
	}

	// Update the user's password and set the status as active
	user.Password = newuser.Password
	user.Status = true

	// Save the changes to the database
	_, err = client.UserMaster.
		Update().
		Where(usermaster.EmployeeIDEQ(user.EmployeeID)).
		SetPassword(user.Password).
		SetStatus(user.Status).
		Save(ctx)

	if err != nil {
		return nil, errors.New("Failed to update user")
	}

	return user, nil
}
