package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"

	//"github.com/twilio/twilio-go"

	//"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	_ "github.com/lib/pq"

	//auth "entdemo/authentication"
	"recruit/authentication"
	"recruit/ent"
	"recruit/ent/login"
	"recruit/ent/usermaster"
	"recruit/start"

	//ent "recruit/ent"
	_ "github.com/lib/pq"
	"recruit/handlers"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	//fmt.Println("hi")
	//client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=secret sslmode=disable")
	client, err := ent.Open("postgres", "host=172.28.12.202 port=2000 user=postgres dbname=recruitment password=secret sslmode=disable")
	//client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Examinations password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//ent.EntClient = client
	initRoutes(client)
	//sendSMS("Dear Customer, OTP for IBC verfication is 234567, please do not share it with anyone - INDPOST", "9884352152", "1007344609998507114", "1001081725895192800", "IBC")
	//sendSMS(msg,mobile,templateid,entityid,appname);

}

func initRoutes(client *ent.Client) {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "PUT", "GET", "OPTIONS"}
	config.MaxAge = 48 * time.Hour
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Cache-Control"}

	r.Use(cors.New(config))
	//fmt.Fprintln(r)
	//smsConfig = gobudgetsms.SetConfig(BsmsUsername, BsmsUserId, BsmsHandle, "", 1, 0, 0)

	//public := r.Group("/rect")
	r.Use(cors.New(config))
	r.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Recruitment APIs"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})

		//public := r.Group("/api")

	})
	//r.POST("/rect/sendsms", generateOTPAndSendSMS)
	//UpdateUserOTPandSendSMS
	r.POST("/rect/sendsms", GenerateOTPAndSendSMSAndSave(client))
	admin := r.Group("/api/admin")
	admin.Use(validateToken)
	admin.Use(totopverify)
	//admin.Use(gettoken)

	r.GET("/api/admin/gettoken/:id", gettoken(client))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/rect/exam", handlers.CreateExam(client))
	r.GET("/rect/exam/:id", handlers.GetExamID(client))
	r.GET("/rect/exams", getExams(client))
	r.POST("/rect/examupdate/:id", updateexam(client))
	r.GET("/rect/exams/ListExamIDsandNames", getExamIDsandNames(client))

	r.POST("/rect/notification", CreateNotification(client))
	r.GET("/rect/notification/:id", GetNotificationID(client))
	r.POST("/rect/notifyupdate/:id", updateNotification(client))
	r.GET("/rect/notifications", getNotifications(client))

	r.POST("/rect/center", CreateCenter(client))
	r.GET("/rect/center/:id", GetCenterID(client))
	r.POST("/rect/center/:id", updateCenter(client))
	r.GET("/rect/centers", getCenters(client))

	r.POST("/rect/profile", CreateUser(client))
	r.GET("/rect/profile/:id", GetUserID(client))
	r.POST("/rect/profile/:id", updateUser(client))
	r.GET("/rect/profiles", getUsers(client))

	r.GET("/rect/nodalofficers", getNodalOfficers(client))
	r.POST("/rect/nodalofficer", CreateNodalOfficer(client))
	r.GET("/rect/nodalofficer/:id", GetNodalOfficerID(client))
	r.POST("/rect/nodalofficer/:id", UpdateNodalOfficer(client))

	r.GET("/rect/examcalendars", getExamCalendars(client))
	r.GET("/rect/examcalendars/:id", GetExamCalendarID(client))
	r.POST("/rect/examcalendars", CreateExamCalendar(client))
	//r.GET("/rect/examcalendars/Details/:id", getExamCalendarsWithDetails(client))

	r.GET("/rect/vacancyyears", getVacancyYears(client))
	r.POST("/rect/vacancyyears", CreateVacarncyYears(client))
	r.GET("/rect/vacancyyears/:id", GetVacancyYearID(client))
	//r.POST("/rect/vacancyyears/:id", UpdateVacancyYearID(client))
	r.GET("/rect/totp", topgenerate)

	r.GET("/rect/exampapers/:id", GetExamPapersByID(client))
	r.GET("/rect/exampapers/byexamcode/:id", GetExamPaperTypesWithPaperCode(client))
	r.GET("/rect/exampapers", getExamPapers(client))
	r.POST("/rect/exampapers", InsertExamPapers(client))

	r.GET("/rect/exampapertypes/:id", GetExamPaperTypesByID(client))
	r.GET("/rect/exampapertypes/bypapercode/:id", GetExamPaperTypesWithPaperCode(client))
	r.GET("/rect/exampapertypes", getAllExamPaperTypes(client))
	r.POST("/rect/exampapertypes", InsertExamPaperTypes(client))

	r.POST("/rect/employees", CreateEmployee(client))
	r.GET("/rect/employees", getAllEmployees(client))
	r.POST("/rect/employeeprofile/verify/:id", updateemployeeverifydetails(client))
	r.GET("rect/employees/:id", GetEmployeesByID(client))

	r.GET("/rect/employeecategories", getAllEmployeeCategories(client))
	r.POST("/rect/employeecategories", CreateEmployeeCategories(client))

	r.GET("/rect/employeedesignation", getAllEmployeeDesignations(client))
	r.POST("/rect/employeedesignation", CreateEmployeeDesignations(client))

	r.GET("/rect/employeedposts", getAllEmployeePosts(client))
	r.GET("/rect/employeedposts/:id", GetEmpPostsByID(client))
	r.POST("/rect/employeedposts", CreateEmployeePost(client))

	r.GET("/rect/employeedisbilitytypes", getAllDisabilityTypes(client))
	r.GET("/rect/employeedisbilitytypes/:id", GetDisabilitiesByID(client))
	r.POST("/rect/employeedisbilitytypes", CreateDisabilityTypes(client))

	r.POST("/rect/circles", CreateCircleMasters(client))
	r.GET("/rect/circles", getAllCircles(client))
	r.GET("/rect/circles/:id", GetCircleByID(client))

	r.POST("/rect/regions", CreateRegionMasters(client))
	r.GET("/rect/regions", getAllRegions(client))
	r.GET("/rect/regions/:id", GetRegionsByID(client))
	r.GET("/rect/regions/bycirclecode/:id", GetRegionsByCircleCode(client))

	r.POST("/rect/divisions", CreateDivisionMasters(client))
	r.GET("/rect/divisions", getAllDivisions(client))
	r.GET("/rect/divisions/:id", GetDivisionsByID(client))
	r.GET("/rect/divisions/byregioncode/:id", GetDivisionsByRegionCode(client))

	r.POST("/rect/facilities", CreateFacilityMasters(client))
	r.GET("/rect/facilities", getAllFacilities(client))
	r.GET("/rect/facilities/:id", GetFacilityByID(client))
	r.GET("/rect/facilities/bycirclecode/:id", GetFacilitiesByCircleCode(client))
	r.GET("/rect/facilities/bydivisioncode/:id", GetFacilitiesByDivisionCode(client))
	r.GET("/rect/facilities/byregioncode/:id", GetFacilitiesByRegionCode(client))
	r.GET("/rect/facilities/byfacilityofficeid/:id", GetFacilityDetailsByFacilityOfficeID(client))

	r.POST("/rect/directorateusers", InsertDirectorateUsers(client))
	r.GET("/rect/directorateusers", getAllDusers(client))
	r.GET("/rect/directorateusers/:id", GetDirectorateUsersByID(client))
	r.GET("/rect/directorateusers/byemployeeid/:id", GetDusersWithEmpID(client))

	r.POST("/rect/eligibilities", InsertEligibilityCriteria(client))
	r.GET("/rect/eligibilities", getAllEligibilitiyCriteria(client))
	r.GET("/rect/eligibilities/:id", GetEligibilityByID(client))
	r.GET("/rect/exams/witheligibilities/byexamcode/:id", getExamWithEligibility(client))

	r.POST("/rect/admin/login/verify", VerifyLogin(client))
	r.GET("/rect/secret/admin/login/all", getAllAdminLogins(client))

	r.POST("/rect/admin/roles", InsertRoles(client))
	r.GET("/rect/admin/roles/:id", GetRolesByID(client))
	r.GET("/rect/admin/roles", GetAllRoles(client))

	r.GET("/rect/users/checkbyusername/:id", CheckAdminUserByUserName(client))
	r.GET("/rect/users/checkbyempid/:id", CheckUserByEmpId(client))
	r.GET("/rect/users/insertuser/:id", CreateUserfromEmpMaster(client))
	r.GET("/rect/users/getuserbyempid/:id", GetUsersByEmpId(client))
	r.POST("/rect/users/login", VerifyUserLogin(client))
	r.POST("/rect/users/new/submit", FirstTimeUserCreation(client))
	r.POST("/rect/adminusers/verifyuser", VerifyAdminUserLogin(client))
	r.POST("/rect/adminusers/verifylogin", VerifyAdminLogin(client))
	r.POST("/rect/users/new/update", UpdateFirstTimeUserDetails(client))
	r.PUT("rect/users/updatepassword/:id", UpdateUserPassword(client))

	r.POST("/rect/ipexams/applications/submit", CreateNewIPApplications(client))
	r.GET("/rect/ipexams/applications/getbyempid/:id", GetIPApplicationsByEmpId(client))
	r.PUT("/rect/ipexams/applications/verify", VerifyApplication(client))
	r.GET("/rect/employees/search/byempid/:id", GetEmployeeDetailsByEmpId(client))
	r.PUT("/rect/ipexams/noverify/:id", UpdateNodalRecommendationsIPByEmpId(client))
	r.GET("/rect/ipexams/getallcapending", getAllCAPendingVerifications(client))
	r.GET("/rect/ipexams/getallcaverified", getAllCAVerified(client))
	r.GET("/rect/ipexams/getallcaverified/:id", GetCAVerifiedDetailsByEmpId(client))
	r.GET("/rect/ipexams/getallcapending/:id", GetCAPendingDetailsByEmpId(client))
	r.GET("/rect/ipexams/caprevremarks/:id", GetCAPendingOldRemarksByEmpId(client))
	r.GET("/rect/ipexams/recommendations/:id", GetIPExamRecommendationsByEmpId(client))
	r.GET("/rect/ipexams/getallnaverified", getAllNAVerified(client))
	r.GET("/rect/ipexams/getallnaverified/:id", GetNAVerifiedDetailsByEmpId(client))

	r.POST("/rect/psexams/applications/submit", CreateNewPSApplications(client))
	r.GET("/rect/psexams/applications/getbyempid/:id", GetPSApplicationsByEmpId(client))

	r.GET("/totp", topgenerate)

	r.POST("/totopverify", totopverify)
	//r.GET("/gettoken/:id", gettoken(client))

	//r.GET("/verify", verifytoken)
	//public.POST("/token", t)
	r.POST("/refreshtoken", validateRefreshToken)
	//r.POST("/token", validateRefreshToken)
	//r.GET("/verify", validateToken)

	r.Run(":8080")
}

func totopverify(c *gin.Context) {
	passcode := c.GetHeader("totp")

	//passcode := c.Request.Header["totp"][0]
	if passcode != "" {
		//c.String(http.StatusUnauthorized, "unauthorised")

		//valid := totp.Validate(passcode, key.Secret())

		valid := totp.Validate(passcode, "MFSWCNBRMVSGGY3FMZTGCZBQGY2WMYZYME4WIZDGGY3DMMJYGUZWMMRZG43TMMBYGVRGKNTFGAYGMMRTGVRWMZDDHAZDAZTGGM4WC")
		if valid {
			//c.String(http.StatusOK, "Valid")
			//os.Exit(0)
		} else {
			c.String(http.StatusUnauthorized, "unauthorised")
			c.Abort()
			return
			//os.Exit(1)
		}

	} else {

		c.String(http.StatusUnauthorized, "unauthorised")
		c.Abort()
		return
	}

	c.Next()

}

func topgenerate(c *gin.Context) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "gateway.cept.gov.in",
		AccountName: "Vams",
		Secret:      []byte("aea41edcceffad065fc8a9ddf6661853f29776085be6e00f235cfdc820ff39a"),
		Period:      90,
		Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"key": key.Secret()})

}

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	UID   string `json:"uid"`
	Uname string `json:"uname"` // struct inside a struct is what is to be implemented if you want all values in single struct
	jwt.StandardClaims
}

func t(gctx *gin.Context) {

	token, err := createToken()
	if err != nil {
		fmt.Println("Couldn't create token error")

	}
	refreshtoken, err := refreshToken()

	if err != nil {
		fmt.Println("Couldn't create token error")

	}

	gctx.JSON(http.StatusOK, gin.H{"token": token, "refreshtoken": refreshtoken})

}

func refreshToken() (refreshtoken string, err error) {
	//Need to make these values dynamic
	var userID uint = 4
	var expireMinutes int = 10
	var secret string = "sadf"

	exp := time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		ID:    userID,
		UID:   uid,
		Uname: "ram",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshtoken, err = jwtToken.SignedString([]byte(secret))

	return

	/*if err != nil {

		fmt.Println("error in creation of token", err.Error())
	}*/

	//gctx.JSON(http.StatusOK, gin.H{"Token": token})
}

func validateRefreshToken(gctx *gin.Context) {

	tokenString := authentication.ExtractToken(gctx)

	claims, err := ParseToken(tokenString, "sadf")
	if err != nil {
		gctx.JSON(http.StatusOK, gin.H{"refreshTokenVerified": "False"})

	}

	if claims != nil {
		if claims.Uname == "ram" {
			//gctx.JSON(http.StatusOK, gin.H{"Username": claims.Uname, "Verified": "True"})

			token, err := createToken()
			if err != nil {
				fmt.Println("Couldn't create token error")

			}
			refreshtoken, err := refreshToken()

			if err != nil {
				fmt.Println("Couldn't create token error")

			}

			fmt.Println("Refresh Token validated.. and token generated")

			gctx.JSON(http.StatusOK, gin.H{"token": token, "refreshtoken": refreshtoken})

		}
	}

	//gctx.JSON(http.StatusOK, gin.H{"Username": "", "Verified": "False"})
}

func validateToken(gctx *gin.Context) {
	tokenString := authentication.ExtractToken(gctx)

	claims, err := ParseToken(tokenString, "faskshfdkljadhfkj")

	fmt.Println(claims)
	if err != nil {
		gctx.String(http.StatusUnauthorized, "unauthorised")
		gctx.Abort()
		return
		//gctx.JSON(http.StatusOK, gin.H{"Username": "", "Verified": "False"})

	}
	if claims.Uname != "ram" {
		gctx.String(http.StatusUnauthorized, "unauthorised")
		gctx.Abort()
		return

		//gctx.JSON(http.StatusOK, gin.H{"Username": claims.Uname, "Verified": "True"})
	}

	gctx.Next()

	/*if claims != nil {
		if claims.Uname == "ram" {

			//gctx.JSON(http.StatusOK, gin.H{"Username": claims.Uname, "Verified": "True"})
		}
	 } */

	//gctx.JSON(http.StatusOK, gin.H{"Username": "", "Verified": "False"})
}

func ParseToken(tokenString, secret string) (
	claims *JwtCustomClaims,
	err error,
) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func createToken() (token string, err error) {
	//Need to make it dynamic
	var userID uint = 4
	var expireMinutes int = 5
	var secret string = "faskshfdkljadhfkj"

	exp := time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		ID:    userID,
		UID:   uid,
		Uname: "ram",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))

	return

	/*if err != nil {

		fmt.Println("error in creation of token", err.Error())
	}*/

	//gctx.JSON(http.StatusOK, gin.H{"Token": token})
}

/*if err != nil {

	fmt.Println("error in creation of token", err.Error())
}*/

//gctx.JSON(http.StatusOK, gin.H{"Token": token})

/*
func verifytoken(gctx *gin.Context) {
	user_id, err := authentication.ExtractTokenID(gctx)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"USERID extracted": user_id})
}

func GenerateToken(gctx *gin.Context) {

	toekn, err := authentication.GenerateToken(233)
	if err != nil {
		fmt.Println("Error in generation of token")
	}
	fmt.Println("Toekn:", toekn)

	gctx.JSON(http.StatusOK, gin.H{"Token": toekn})
}*/

func getExams(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exams, err := start.QueryExam(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exams})

	}

	return gin.HandlerFunc(fn)

}

func getExamIDsandNames(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exams, err := start.QueryExamIDAndNames(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exams})

	}

	return gin.HandlerFunc(fn)

}
func updateexam(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExam := new(ent.Exam)

		id := gctx.Param("id")

		examID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newExam); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		exam, err := start.UpdateExam(client, int32(examID), newExam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exam})

	}

	return gin.HandlerFunc(fn)
}

func CreateUser(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newUser := new(ent.User)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser, err := start.CreateUser(client, newUser)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the User")

	}
	return gin.HandlerFunc(fn)
}

func GetUserID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		//UserID, _ := strconv.ParseInt(id, 10, 32)

		center, err := start.QueryUserID(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": center})

	}

	return gin.HandlerFunc(fn)
}

func updateUser(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.User)

		id := gctx.Param("id")

		//CenterID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := start.UpdateUser(client, id, newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": user})

	}

	return gin.HandlerFunc(fn)
}

func getUsers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		users, err := start.QueryUser(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": users})

	}

	return gin.HandlerFunc(fn)

}

func CreateCenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newCenter := new(ent.Center)

		if err := gctx.ShouldBindJSON(&newCenter); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCenter, err := start.CreateCenter(client, newCenter)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Center")

	}
	return gin.HandlerFunc(fn)
}

func GetCenterID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		CenterID, _ := strconv.ParseInt(id, 10, 32)

		center, err := start.QueryCenterID(ctx, client, int32(CenterID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": center})

	}

	return gin.HandlerFunc(fn)
}

func updateCenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newCenter := new(ent.Center)

		id := gctx.Param("id")

		CenterID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newCenter); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		center, err := start.UpdateCenter(client, int32(CenterID), newCenter)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": center})

	}

	return gin.HandlerFunc(fn)
}

func getCenters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		centers, err := start.QueryCenter(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": centers})

	}

	return gin.HandlerFunc(fn)

}

func CreateNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newNotification := new(ent.Notification)

		if err := gctx.ShouldBindJSON(&newNotification); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newNotification, err := start.CreateNotification(client, newNotification)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Notification")

	}
	return gin.HandlerFunc(fn)
}

func GetNotificationID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		NotificationID, _ := strconv.ParseInt(id, 10, 32)

		notification, err := start.QueryNotificationID(ctx, client, int32(NotificationID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": notification})

	}

	return gin.HandlerFunc(fn)
}

func updateNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newNotification := new(ent.Notification)

		id := gctx.Param("id")

		NotificationID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newNotification); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		exam, err := start.UpdateNotification(client, int32(NotificationID), newNotification)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exam})

	}

	return gin.HandlerFunc(fn)
}

func getNotifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		notfications, err := start.QueryNotification(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": notfications})

	}

	return gin.HandlerFunc(fn)

}
func getNodalOfficers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		NodalOfficers, err := start.QueryNodalOfficer(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": NodalOfficers})

	}

	return gin.HandlerFunc(fn)

}
func CreateNodalOfficer(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newNodalOfficer := new(ent.NodalOfficer)

		if err := gctx.ShouldBindJSON(&newNodalOfficer); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newNodalOfficer, err := start.CreateNodalOfficer(client, newNodalOfficer)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Nodal Officer")

	}
	return gin.HandlerFunc(fn)
}

func GetNodalOfficerID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		NodalOfficerID, _ := strconv.ParseInt(id, 10, 32)

		nodalofficer, err := start.QueryNodalOfficerID(ctx, client, int32(NodalOfficerID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": nodalofficer})

	}

	return gin.HandlerFunc(fn)
}
func UpdateNodalOfficer(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newNodalOfficer := new(ent.NodalOfficer)
		id := gctx.Param("id")
		NodalOfficerID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newNodalOfficer); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		nodalofficer, err := start.UpdateNodalOfficer(client, int32(NodalOfficerID), newNodalOfficer)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": nodalofficer})

	}

	return gin.HandlerFunc(fn)
}

// For Exam Calendars ...!
func CreateExamCalendar(client *ent.Client) gin.HandlerFunc {

	fmt.Println("came inside exam calenders")
	fn := func(gctx *gin.Context) {
		newExamCalendar := new(ent.ExamCalendar)
		if err := gctx.ShouldBindJSON(&newExamCalendar); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newExamCalendar, err := start.CreateExamCalendar(client, newExamCalendar)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam Calendar")
	}
	return gin.HandlerFunc(fn)
}

func GetExamCalendarID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		calendarID, _ := strconv.ParseInt(id, 10, 32)

		calendars, err := start.QueryExamCalendarID(ctx, client, int32(calendarID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": calendars})

	}

	return gin.HandlerFunc(fn)
}

func getExamCalendars(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		examcalendars, err := start.QueryExamCalendars(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcalendars})

	}

	return gin.HandlerFunc(fn)

}

/*func getExamCalendarsWithDetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		ExamCode, _ := strconv.ParseInt(id, 10, 32)
		examcals, err := start.QueryExamCalendarsWithVacancyAndPapers(ctx, client, int32(ExamCode))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcals})
	}
	return gin.HandlerFunc(fn)
}*/

func CreateExamCalendarData(client *ent.Client) gin.HandlerFunc {
	// Parse the request body and extract the necessary data
	fn := func(gctx *gin.Context) {
		//var requestData ExamCalendar
		newExamCalendar := new(ent.ExamCalendar)
		if err := gctx.ShouldBindJSON(&newExamCalendar); err != nil {
			// Handle invalid request data error
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// Start a new transaction
		newExamCalendar, err := start.CreateExamCalendar(client, newExamCalendar)
		if err != nil {
			// Handle transaction error
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Notification")
	}
	return gin.HandlerFunc(fn)
}

// For Vacancy Years ...!
func CreateVacarncyYears(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newVacancyYear := new(ent.VacancyYear)

		if err := gctx.ShouldBindJSON(&newVacancyYear); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newVacancyYear, err := start.CreateVacancyYears(client, newVacancyYear)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Vacancy Year")

	}
	return gin.HandlerFunc(fn)
}

func GetVacancyYearID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		VacancyYearID, _ := strconv.ParseInt(id, 10, 32)

		vys, err := start.QueryVacancyYearID(ctx, client, int32(VacancyYearID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": vys})

	}

	return gin.HandlerFunc(fn)
}

func UpdateVacancyYearID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newVacYr := new(ent.VacancyYear)
		id := gctx.Param("id")
		VacancyYearID, _ := strconv.ParseInt(id, 10, 32)
		if err := gctx.ShouldBindJSON(&newVacYr); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vcy, err := start.UpdateVacancyYearID(client, int32(VacancyYearID), newVacYr)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": vcy})
	}
	return gin.HandlerFunc(fn)
}

func getVacancyYears(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		vcs, err := start.QueryVacancyYear(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": vcs})

	}

	return gin.HandlerFunc(fn)

}

// For ExamPapers
func InsertExamPapers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExamPapers := new(ent.ExamPapers)
		if err := gctx.ShouldBindJSON(&newExamPapers); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newExamPapers, err := start.CreatePapers(client, newExamPapers)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam Papers")
	}
	return gin.HandlerFunc(fn)
}

func GetExamPapersByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		ExamPaperID, _ := strconv.ParseInt(id, 10, 32)
		exampapers, err := start.QueryExamPaperByID(ctx, client, int32(ExamPaperID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampapers})
	}
	return gin.HandlerFunc(fn)
}

func getAllExamPapers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exampapers, err := start.QueryExamPapers(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampapers})

	}

	return gin.HandlerFunc(fn)

}

// Exam Papers
func getExamPapers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exampaps, err := start.QueryExamPapers(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampaps})

	}

	return gin.HandlerFunc(fn)

}

// QueryExamPapersByExamCode
// QueryCircleMasterWithRegions
func GetExamPapersWithExamCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		ExamPaperID, _ := strconv.ParseInt(id, 10, 32)
		exampapers, err := start.QueryExamPapersByExamCode(ctx, client, int32(ExamPaperID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get ExamPapers by ExamCode and retrieve only ID and PaperDescription.

		// Create an array to store the PaperID and PaperDescription.
		var paperData []map[string]interface{}
		for _, paper := range exampapers {
			paperData = append(paperData, map[string]interface{}{
				"PaperID":          paper.ID,
				"PaperDescription": paper.PaperDescription,
			})
		}
		gctx.JSON(http.StatusOK, gin.H{"data": paperData})
	}
	return gin.HandlerFunc(fn)
}

// For Employees
func CreateEmployee(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmployees := new(ent.Employees)
		if err := gctx.ShouldBindJSON(&newEmployees); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newEmployees, err := start.CreateEmployeeProfile(client, newEmployees)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Profile")
	}
	return gin.HandlerFunc(fn)
}

func getAllEmployees(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		newEmployees, err := start.QueryEmployees(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployees})
	}

	return gin.HandlerFunc(fn)

}

func GetEmployeesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		EmpID, _ := strconv.ParseInt(id, 10, 32)
		emps, err := start.QueryEmployeesWithID(ctx, client, int32(EmpID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": emps})
	}
	return gin.HandlerFunc(fn)
}
func updateemployeeverifydetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmp := new(ent.Employees)
		id := gctx.Param("id")
		empID, _ := strconv.ParseInt(id, 10, 32)
		if err := gctx.ShouldBindJSON(&newEmp); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		empupdated, err := start.UpdateVerificationDetails(client, int32(empID), newEmp)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": empupdated})

	}

	return gin.HandlerFunc(fn)
}

//	r.GET("/rect/employeecategories/", getAllEmployeeCategories(client))
//
// r.POST("/rect/exampapers/", CreateEmployeeCategories(client)
func CreateEmployeeCategories(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmployeecategories := new(ent.EmployeeCategory)
		if err := gctx.ShouldBindJSON(&newEmployeecategories); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newEmployeecategories, err := start.CreateEmployeeCategory(client, newEmployeecategories)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Category")
	}
	return gin.HandlerFunc(fn)
}

func getAllEmployeeCategories(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		newEmployeCategories, err := start.QueryEmployeeCategories(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployeCategories})
	}

	return gin.HandlerFunc(fn)

}

// Employee Designation
func CreateEmployeeDesignations(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmployeeDesignations := new(ent.EmployeeDesignation)
		if err := gctx.ShouldBindJSON(&newEmployeeDesignations); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newEmployeeDesignations, err := start.CreateEmployeeDesignation(client, newEmployeeDesignations)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Category")
	}
	return gin.HandlerFunc(fn)
}

func getAllEmployeeDesignations(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		newEmployeeDesignations, err := start.QueryEmployeeDesignations(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployeeDesignations})
	}

	return gin.HandlerFunc(fn)

}

// QueryFacilities

/*func getAllFacilities(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		newEmployeCategories, err := start.QueryFacilities(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployeCategories})
	}

	return gin.HandlerFunc(fn)

}
*/
// GetCircles bY ID

func GetCircleByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		CircleID, _ := strconv.ParseInt(id, 10, 32)

		circles, err := start.QueryCircleMasterByID(ctx, client, int32(CircleID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)
}

// For Circle Master
func CreateCircleMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newCircles := new(ent.CircleMaster)
		if err := gctx.ShouldBindJSON(&newCircles); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newCircles, err := start.CreateCircleMaster(client, newCircles)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Circle Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllCircles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryCircleMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}

// For Region Master
func CreateRegionMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newRegions := new(ent.RegionMaster)
		if err := gctx.ShouldBindJSON(&newRegions); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newRegions, err := start.CreateRegionMaster(client, newRegions)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Regions Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllRegions(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryRegionMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}

// QueryCircleMasterWithRegions
func GetRegionsByCircleCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		CircleID, _ := strconv.ParseInt(id, 10, 32)
		circles, err := start.QueryRegionMasterByCircleCode(ctx, client, int32(CircleID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}
func GetRegionsByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		CircleID, _ := strconv.ParseInt(id, 10, 32)

		regions, err := start.QueryRegionMasterByID(ctx, client, int32(CircleID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": regions})

	}

	return gin.HandlerFunc(fn)
}

// For Division Master
func CreateDivisionMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newDivisions := new(ent.DivisionMaster)
		if err := gctx.ShouldBindJSON(&newDivisions); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newDivisions, err := start.CreateDivisionMaster(client, newDivisions)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Divisions Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllDivisions(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		divisions, err := start.QueryDivisionMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": divisions})

	}
	return gin.HandlerFunc(fn)
}

// QueryCircleMasterWithRegions
func GetDivisionsByRegionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		DivisionID, _ := strconv.ParseInt(id, 10, 32)
		divisions, err := start.QueryDivisionMasterByRegionCode(ctx, client, int32(DivisionID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": divisions})
	}
	return gin.HandlerFunc(fn)
}
func GetDivisionsByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		DivID, _ := strconv.ParseInt(id, 10, 32)

		divisions, err := start.QueryDivisionMasterByID(ctx, client, int32(DivID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": divisions})

	}

	return gin.HandlerFunc(fn)
}

// Facilities
func CreateFacilityMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newfac := new(ent.Facility)
		if err := gctx.ShouldBindJSON(&newfac); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newfac, err := start.CreateFacility(client, newfac)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Facility Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllFacilities(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		facs, err := start.QueryFacilitiesMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})

	}
	return gin.HandlerFunc(fn)
}

func GetFacilitiesByDivisionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryFacilityMasterByDivisionCode(ctx, client, int32(FacID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}
func GetFacilityByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		facID, _ := strconv.ParseInt(id, 10, 32)

		facs, err := start.QueryFacilityMasterByID(ctx, client, int32(facID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})

	}

	return gin.HandlerFunc(fn)
}

func GetFacilitiesByRegionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryFacilityMasterByRegionCode(ctx, client, int32(FacID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}

func GetFacilitiesByCircleCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryFacilityMasterByCircleCode(ctx, client, int32(FacID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}

// / Employee posts
func CreateEmployeePost(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newposts := new(ent.EmployeePosts)
		if err := gctx.ShouldBindJSON(&newposts); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newposts, err := start.CreateEmployeePosts(client, newposts)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Posts")
	}
	return gin.HandlerFunc(fn)
}

func getAllEmployeePosts(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		posts, err := start.QueryEmployeePosts(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": posts})

	}
	return gin.HandlerFunc(fn)
}

func GetEmpPostsByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		postID, _ := strconv.ParseInt(id, 10, 32)

		posts, err := start.QueryEmployeePostsID(ctx, client, int32(postID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": posts})

	}

	return gin.HandlerFunc(fn)
}

// / Disabilities
func CreateDisabilityTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newdisability := new(ent.Disability)
		if err := gctx.ShouldBindJSON(&newdisability); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newdisability, err := start.CreateEmployeeDisability(client, newdisability)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Disability Types")
	}
	return gin.HandlerFunc(fn)
}

func getAllDisabilityTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		disability, err := start.QueryEmployeeDisabilities(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": disability})

	}
	return gin.HandlerFunc(fn)
}

func GetDisabilitiesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		disabilityID, _ := strconv.ParseInt(id, 10, 32)

		disability, err := start.QueryEmployeeDisabilityID(ctx, client, int32(disabilityID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": disability})

	}

	return gin.HandlerFunc(fn)
}

// Directorate Users

func InsertDirectorateUsers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		Dusers := new(ent.DirectorateUsers)
		if err := gctx.ShouldBindJSON(&Dusers); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Dusers, err := start.CreateDirectorateUsers(client, Dusers)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the DirectorateUsers")
	}
	return gin.HandlerFunc(fn)
}

func GetDirectorateUsersByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		DuserID, _ := strconv.ParseInt(id, 10, 32)
		Dusers, err := start.QueryDUsersByID(ctx, client, int32(DuserID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})
	}
	return gin.HandlerFunc(fn)
}

func getAllDusers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		Dusers, err := start.QueryeDirectorateUsers(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

}

func GetDusersWithEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		DuserID, _ := strconv.ParseInt(id, 10, 32)
		Dusers, err := start.QueryeDirectorateUsersyEmpId(ctx, client, int32(DuserID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get Dusers by Emp ID and retrieve only details

		// Create an array to store the PaperID and PaperDescription.
		var DuserData []map[string]interface{}
		for _, user := range Dusers {
			DuserData = append(DuserData, map[string]interface{}{
				"EmpID":        user.EmployeedID,
				"EmployeeName": user.EmployeeName,
				"Role":         user.Role,
				"MobileNumber": user.MobileNumber,
				"EmailId":      user.EmailId,
			})
		}
		gctx.JSON(http.StatusOK, gin.H{"data": DuserData})

	}
	return gin.HandlerFunc(fn)
}

// Eligibility Master

func InsertEligibilityCriteria(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		eligibilities := new(ent.EligibilityMaster)
		if err := gctx.ShouldBindJSON(&eligibilities); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		eligibilities, err := start.CreateEligibilities(client, eligibilities)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Eligibility criteria")
	}
	return gin.HandlerFunc(fn)
}

func GetEligibilityByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		eliID, _ := strconv.ParseInt(id, 10, 32)
		elis, err := start.QueryEligibilitiyMasterByID(ctx, client, int32(eliID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": elis})
	}
	return gin.HandlerFunc(fn)
}

func getAllEligibilitiyCriteria(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		Dusers, err := start.QueryEligibilitiyMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

}

func getExamWithEligibility(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		examcode, _ := strconv.ParseInt(id, 10, 32)
		exams, err := start.QueryEligibilityWithPapers(ctx, client, int32(examcode))

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exams})

	}
	return gin.HandlerFunc(fn)
}

/*func getLoginDetails(client *ent.Client) gin.HandlerFunc {
fn := func(gctx *gin.Context) {
	ctx := context.Background()
	id := gctx.Param("id")
	//var examID int32
	empid, _ := strconv.ParseInt(id, 10, 32)
	logins, err := start.QueryLoginByEmpID(ctx, client, int32(empid))

	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"data": logins})

}
}*/

//createtokenfrom db /*

func createTokenfromdb(ctx context.Context, client *ent.Client, empid int32) (string, error) {
	// Retrieve dynamic values from the database using the Ent client

	userIDFromDB := uint(0)
	expireMinutesFromDB := 0
	secretFromDB := ""

	// Example of retrieving values from the database using Ent
	// Replace with your own implementation based on your Ent schema
	userEnt, err := client.Login.Query().
		Where(login.EmployeedIDEQ(empid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			// Handle the case when no entity is found for the given condition
			log.Println("No login entry found for the specified employee ID")
			return "", fmt.Errorf("no login entry found for the specified employee ID: %w", err)
		}
		// Handle other error cases
		log.Println("Error getting login details from the database:", err)
		return "", fmt.Errorf("failed to retrieve login details from the database: %w", err)
	}

	// Log the retrieved values
	log.Printf("Retrieved login details for employee ID %d:", empid)
	log.Printf("UserID: %d", userEnt.EmployeedID)
	log.Printf("ExpireMinutes: %d", userEnt.ExpireminsToken)
	log.Printf("Secret: %s", userEnt.LoginID)

	userIDFromDB = uint(userEnt.EmployeedID)
	expireMinutesFromDB = int(userEnt.ExpireminsToken)
	//secretFromDB = userEnt.LoginID.String() // Convert UUID to string
	secretFromDB = userEnt.Username

	exp := time.Now().Add(time.Minute * time.Duration(expireMinutesFromDB)).Unix()
	uid := uuid.New().String()
	claims := &JwtCustomClaims{
		ID:    userIDFromDB,
		UID:   uid,
		Uname: userEnt.Username,
		//Uname: "ram",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secretFromDB))

	log.Println("Generated token:", token)

	return token, err
}

func gettoken(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		empid, _ := strconv.ParseInt(id, 10, 32)
		logins, err := createTokenfromdb(ctx, client, int32(empid))
		if err != nil {
			log.Println("Invalid empid:", err)
			// Handle the error appropriately
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": logins})
	}
}

//For Admin Login

func InsertAdminLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		adminlogins := new(ent.AdminLogin)
		if err := gctx.ShouldBindJSON(&adminlogins); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		adminlogins, err := start.CreateAdminLogin(client, adminlogins)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Admin Login")
	}
	return gin.HandlerFunc(fn)
}

func VerifyLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		logins := new(ent.AdminLogin)
		if err := gctx.ShouldBindJSON(&logins); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logins, err := start.ValidateAdminLoginUser(client, logins)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			//"message": "User " + logins.Username + " verified successfully " + "for the role " + logins.RoleName,

			"data": gin.H{
				"roleUserCode":  logins.RoleUserCode,
				"verifyRemarks": logins.VerifyRemarks,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

func getAllAdminLogins(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		Dusers, err := start.QueryAdminLogin(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

}

// RoleMaster
func InsertRoles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newroles := new(ent.RoleMaster)
		if err := gctx.ShouldBindJSON(&newroles); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newroles, err := start.CreateRoles(client, newroles)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created tnew roles")
	}
	return gin.HandlerFunc(fn)
}

func GetRolesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		eliID, _ := strconv.ParseInt(id, 10, 32)
		elis, err := start.QueryRolesByID(ctx, client, int32(eliID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": elis})
	}
	return gin.HandlerFunc(fn)
}
func GetAllRoles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		Dusers, err := start.QueryRoles(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

}

// / Exam Papertypes
func InsertExamPaperTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExamPapertypes := new(ent.PaperTypes)
		if err := gctx.ShouldBindJSON(&newExamPapertypes); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newExamPapertypes, err := start.CreatePaperType(client, newExamPapertypes)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam Paper Types")
	}
	return gin.HandlerFunc(fn)
}

func GetExamPaperTypesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		ExamPaperTypeID, _ := strconv.ParseInt(id, 10, 32)
		exampapertypes, err := start.QueryExamPaperTypeByID(ctx, client, int32(ExamPaperTypeID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampapertypes})
	}
	return gin.HandlerFunc(fn)
}

func getAllExamPaperTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		papertypes, err := start.QueryExamPaperTypes(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": papertypes})

	}

	return gin.HandlerFunc(fn)

}

// QueryExamPapersByExamCode
// QueryCircleMasterWithRegions
func GetExamPaperTypesWithPaperCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		PaperID, _ := strconv.ParseInt(id, 10, 32)
		papertypes, err := start.QueryExamPaperTypesByPaperCode(ctx, client, int32(PaperID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get ExamPapers by ExamCode and retrieve only ID and PaperDescription.

		// Create an array to store the PaperID and PaperDescription.
		var papertypeData []map[string]interface{}
		for _, paper := range papertypes {
			papertypeData = append(papertypeData, map[string]interface{}{
				"PaperID":          paper.ID,
				"PaperDescription": paper.PaperTypeDescription,
			})
		}
		gctx.JSON(http.StatusOK, gin.H{"data": papertypeData})
	}
	return gin.HandlerFunc(fn)
}

// Send SMS
type SMSRequest struct {
	Msg        string `form:"msg" binding:"required"`
	Phone      string `form:"phone" binding:"required"`
	TemplateID string `form:"templateid" binding:"required"`
	EntityID   string `form:"entityid" binding:"required"`
	AppName    string `form:"appname" binding:"required"`
}

func sendSMS(c *gin.Context) {
	var req SMSRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	url := "https://api.cept.gov.in/sendsms/api/values/sendsms"

	payloadBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode request data"})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS"})
		return
	}
	defer resp.Body.Close()

	// Print the response body for debugging
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}
	fmt.Println(string(responseBody))

	var result string
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

//Query for existing User from user master by user name

func CheckAdminUserByUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//CircleID, _ := strconv.ParseInt(id, 10, 32)
		user, err := start.QueryUserMasterByUserName(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": user})
	}
	return gin.HandlerFunc(fn)
}

// Query by Emp ID

type NewUserResponse struct {
	NewUser bool `json:"NewUser"`
}

func (r NewUserResponse) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`{"NewUser": %v}`, r.NewUser)
	return []byte(str), nil
}

func CheckUserByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		empID, _ := strconv.ParseInt(id, 10, 64)
		user, err := start.QueryUserMasterByEmpId(ctx, client, empID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response := NewUserResponse{
			NewUser: user,
		}
		gctx.JSON(http.StatusOK, gin.H{"data": response})
	}
	return gin.HandlerFunc(fn)
}

//CreateUserByEmpId

func CreateUserfromEmpMaster(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		id := gctx.Param("id")
		empID, _ := strconv.ParseInt(id, 10, 64)

		user, err := start.CreateUserByEmpId(ctx, client, empID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return the specified fields
		/*gctx.JSON(http.StatusOK, gin.H{
			"employeeID":    user.EmployeeID,
			"employeeName":  user.EmployeeName,
			"role":          user.Role,
			"mobile":        user.Mobile,
		})*/
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":         user.EmployeeID,
				"NewPasswordRequest": user.NewPasswordRequest,
				"RoleUserCode":       user.RoleUserCode,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// QueryUsersByEmpId
func GetUsersByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryUsersByEmpId(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

/*func updateUserWithPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newemp := new(ent.UserMaster)

		id := gctx.Param("id")

		empid, _ := strconv.ParseInt(id, 10, 64)

		if err := gctx.ShouldBindJSON(&newemp); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		emps, err := start.UpdateUserByEmpId(client, empid, newemp)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": emps})
	}
	return gin.HandlerFunc(fn)
}
*/

func UpdateUser(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.UserMaster)

		id := gctx.Param("id")

		userID, _ := strconv.ParseInt(id, 10, 64)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := start.UpdateUserByEmpID(client, int64(userID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":     user.EmployeeID,
				"PasswordStatus": "Password updated Successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// Verify User Login
func VerifyUserLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		// Parse the request body into an AdminLogin entity
		login := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&login); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the login credentials using the ValidateLoginUser function
		user, err := start.ValidateLoginUser(client, &ent.UserMaster{
			UserName: login.UserName,
			Password: login.Password,
		})
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"RoleUserCode": user.RoleUserCode,
				"UserName":     user.UserName,
				"LoginStatus":  "Verified successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
}

func UpdateUserPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.UserMaster)

		id := gctx.Param("id")

		userID, _ := strconv.ParseInt(id, 10, 64)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := start.UpdateUserPasswordByEmpID(client, int64(userID), newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":     user.EmployeeID,
				"PasswordStatus": "Password updated Successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
}
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

	//fmt.Println("Response body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	response := string(body)
	if !strings.Contains(response, "SMS Pushed to NIC Successfully") {
		return fmt.Errorf("failed to send SMS")
	}

	return nil
}
func GenerateOTP() int32 {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	otp := rand.Intn(max-min+1) + min
	return int32(otp)
}

func generateOTPAndSendSMS(c *gin.Context) {
	// Declare a struct to bind the request JSON data
	var requestData struct {
		EmpID int64 `json:"emp_id"`
	}

	// Bind the request JSON to the struct
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Retrieve the employee ID from the request data
	//	empID := requestData.EmpID
	//empID := requestData.EmpID
	// Generate OTP
	otp := GenerateOTP()
	stringOTP := strconv.Itoa(int(otp))

	// Retrieve the user's mobile number from the database or any other source
	mobileNumber := "9884352152"
	//mobileNumber := start.getUserMobileNumber(empID)
	if mobileNumber == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Construct the SMS message
	msg := "Dear Customer, OTP for IBC verification is " + stringOTP + ", please do not share it with anyone - INDPOST"

	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007344609998507114"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	err := sendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS"})
		c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})

	}

	// Update OTP for the user

	//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS"})
}

func updateOTP(client *ent.Client, empID int64, otp int32) (*ent.UserMaster, error) {
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

	userMaster, err = userMaster.Update().
		SetOTP(otp).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update OTP: %v", err)
	}

	return userMaster, nil
}

func GenerateOTPAndSendSMSAndSave(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		newUser := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//user, err := start.CreateUserByEmpId(ctx, client, empID)
		_, message, err := start.SendSMSAndSaveOTP(ctx, client, newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"message": message})
	}
	return gin.HandlerFunc(fn)
}

// Create IP Applications with Circle Preferences ...
func CreateNewIPApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newIPAppplns := new(ent.Exam_Applications_IP)
		if err := gctx.ShouldBindJSON(&newIPAppplns); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newIPAppplns, err := start.CreateIPApplications(client, newIPAppplns)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":         newIPAppplns.EmployeeID,
				"Application Status": newIPAppplns.ApplicationStatus,
				"ApplicationNumber":  newIPAppplns.ApplicationNumber,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// Get IP Applications with EmpID - QueryIPExamApplicationsByEmpID
func GetIPApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		ipapplns, err := start.QueryIPExamApplicationsByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": ipapplns})
	}
	return gin.HandlerFunc(fn)
}

/*func GetIPApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		empid, _ := strconv.ParseInt(id, 10, 64)
		ipapplns, err := start.QueryIPExamApplicationsByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Query the CirclePrefRef edges
		circlePrefs, err := ipapplns.QueryCirclePrefRef().All(ctx)
		if err != nil {
			log.Println("error querying CirclePrefRef edges: ", err)
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve CirclePrefRef edges"})
			return
		}

		// Extract only the desired fields from circlePrefs
		var result struct {
			CirclePrefRef []struct {
				PlacePrefNo    int32  `json:"place_pref_no"`
				PlacePrefValue string `json:"place_pref_value"`
			} `json:"circle_pref_ref"`
		}
		for _, edge := range circlePrefs {
			result.CirclePrefRef = append(result.CirclePrefRef, struct {
				PlacePrefNo    int32  `json:"place_pref_no"`
				PlacePrefValue string `json:"place_pref_value"`
			}{
				PlacePrefNo:    edge.PlacePrefNo,
				PlacePrefValue: edge.PlacePrefValue,
			})
		}

		gctx.JSON(http.StatusOK, gin.H{"data": result})
	}
	return gin.HandlerFunc(fn)
}*/

// PS Group B
// JSON collection with preferences...
func CreateNewPSApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newPSAppplns := new(ent.Exam_Applications_PS)
		if err := gctx.ShouldBindJSON(&newPSAppplns); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newIPAppplns, err := start.CreatePSApplications(client, newPSAppplns)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":         newIPAppplns.EmployeeID,
				"Application Status": newIPAppplns.ApplicationStatus,
				//"ApplicationNumber":  newIPAppplns.ApplicationNumber,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// Get IP Applications with EmpID - QueryPSExamApplicationsByEmpID
func GetPSApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		psapplns, err := start.QueryPSExamApplicationsByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": psapplns})
	}
	return gin.HandlerFunc(fn)
}

//Update Verification details

func VerifyApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newAppln := new(ent.Exam_Applications_IP)

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		application, err := start.UpdateApplicationRemarks(client, newAppln)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":          application.EmployeeID,
				"ApplicationStatus":   application.ApplicationStatus,
				"Application Remarks": application.AppliactionRemarks,
				//	"CAOldRemarks":        application.CAPreviousRemarks,
			},
		})
	}

	return gin.HandlerFunc(fn)
}

// Get Employee Details from Employee Master

func GetEmployeeDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryEmployeeMasterByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Update Nodal recommendations
func UpdateNodalRecommendationsIPByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		id := gctx.Param("id")
		empid, _ := strconv.ParseInt(id, 10, 64)

		// Bind request data to an array of ent.RecommendationsIPApplications
		var newRecommendations []*ent.RecommendationsIPApplications

		// Bind the request JSON to the array of ent.RecommendationsIPApplications
		if err := gctx.ShouldBindJSON(&newRecommendations); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedRecords, err := start.UpdateNodalRecommendationsByEmpID(client, empid, newRecommendations)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Filter the updated records to include only the relevant empID
		filteredRecords := make([]*ent.RecommendationsIPApplications, 0)
		recordMap := make(map[int64]bool) // Map to track processed records

		for _, record := range updatedRecords {
			if record.EmployeeID == empid {
				// Check if the record has already been added
				if _, ok := recordMap[record.ID]; !ok {
					filteredRecords = append(filteredRecords, record)
					recordMap[record.ID] = true // Mark the record as processed
				}
			}
		}

		// Log the filtered records
		//log.Printf("Filtered records for Employee ID %d: %+v", empid, filteredRecords)

		gctx.JSON(http.StatusOK, gin.H{"data": filteredRecords})
	}
	return gin.HandlerFunc(fn)
}

// get ca pending verify records
func getAllCAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryIPApplicationsByCAVerificationsPending(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}

// Get All CA Verified records
func getAllCAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryIPApplicationsByCAVerified(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}

// CA Verified with Emp ID
func GetCAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryIPApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

//CA Pending with EmpID -- QueryIPApplicationsByCAPendingByEmpID

// CA Verified with Emp ID
func GetCAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryIPApplicationsByCAPendingByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Return Previous Remarks
func GetCAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//newAppln := new(ent.Exam_Applications_IP)
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.GetOldCAApplicationRemarksByEmployeeID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": circles})

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":          circles.EmployeeID,
				"ApplicationStatus":   circles.ApplicationStatus,
				"Application Remarks": circles.AppliactionRemarks,
				//	"CAOldRemarks":        application.CAPreviousRemarks,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// Verify AdminLoginUser
func VerifyAdminUserLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		// Parse the request body into an AdminLogin entity
		login := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&login); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the login credentials using the ValidateLoginUser function
		user, err := start.ValidateAdminUserLogin(client, &ent.UserMaster{
			UserName: login.UserName,
			Password: login.Password,
		})
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"RoleUserCode":   user.RoleUserCode,
				"UserName":       user.UserName,
				"PasswordStatus": "Verified successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// ValidateAdminLogin
func VerifyAdminLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		// Parse the request body into an AdminLogin entity
		login := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&login); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the login credentials using the ValidateLoginUser function
		user, err := start.ValidateAdminLogin(client, &ent.UserMaster{
			UserName: login.UserName,
			OTP:      login.OTP,
		})
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"RoleUserCode": user.RoleUserCode,
				"UserName":     user.UserName,
				"LoginStatus":  "Verified successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// GetRecommendationsByEmpID
func GetIPExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, err := start.GetRecommendationsByEmpID(client, empid)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": records})
	}
	return gin.HandlerFunc(fn)
}

// QueryFacilityMasterDetailsByID
func GetFacilityDetailsByFacilityOfficeID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		officeID := gctx.Param("id")

		facility, err := start.QueryFacilityByOfficeID(ctx, client, officeID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{"data": facility})
	}

	return gin.HandlerFunc(fn)
}

// Create User and Send SMS
func FirstTimeUserCreation(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		newUser := new(ent.UserMaster)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := start.NewCreateUserByEmpId(ctx, client, newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":         user.EmployeeID,
				"UserName":           user.UserName,
				"UserCreationStatus": "User created successfully and OTP triggered",
				"RoleUserCode":       user.RoleUserCode,
				//	"CAOldRemarks":        application.CAPreviousRemarks,
			},
		})
	}

	return gin.HandlerFunc(fn)
}

// Update Password of First time User into User Master
func UpdateFirstTimeUserDetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		// Parse the request body into an AdminLogin entity
		login := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&login); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the login credentials using the ValidateLoginUser function
		user, err := start.NewValidateLoginUser(client, &ent.UserMaster{
			UserName: login.UserName,
			OTP:      login.OTP,
			Password: login.Password,
		})
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"RoleUserCode": user.RoleUserCode,
				"UserName":     user.UserName,
				//"PasswordStatus":  "Updated successfully",
				"Message": "User registered successfully and can continue to Login !!",
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// Get All NA Verified records ...
func getAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryIPApplicationsByNAVerified(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}

// Get na verified record by Emp ID
func GetNAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryIPApplicationsByNAVerifiedByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}
