package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/usermaster"
	"strconv"
)

func SendSMSAndSaveOTP(ctx context.Context, client *ent.Client, newUser *ent.UserMaster) (string, string, error) {
	// Check if username is not null
	if newUser.UserName == "" {
		return "", "", errors.New("Username is null")
	}

	// Retrieve the user from the database based on the username
	user, err := client.UserMaster.Query().Where(usermaster.UserNameEQ(newUser.UserName)).Only(ctx)
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Retrieve the mobile number from the retrieved user
	mobileNumber := user.Mobile
	if mobileNumber == "" {
		return "", "", errors.New("User's mobile number not found")
	}

	// Generate OTP
	otp := GenerateOTP()

	// Convert OTP to string
	stringOTP := strconv.Itoa(int(otp))

	// Construct the SMS message
	msg := "Dear Customer, OTP for IBC verification is " + stringOTP + ", please do not share it with anyone - INDPOST"

	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007344609998507114"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	err = sendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Printf("Failed to send SMS: %v", err)
		return "", "", fmt.Errorf("failed to send OTPs")
	}

	// Save the OTP in the UserMaster entity
	updatedUser, err := user.Update().SetOTP(otp).Save(ctx)
	if err != nil {
		log.Printf("Failed to save OTP in UserMaster: %v", err)
		return "", "", fmt.Errorf("failed to save OTP")
	}

	log.Printf("OTP sent and saved successfully. User: %s, OTP: %s", updatedUser.UserName, stringOTP)

	return stringOTP, "OTP sent Successfully.", nil
}
