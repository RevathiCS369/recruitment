package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"

	//"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	_ "github.com/lib/pq"

	//auth "entdemo/authentication"
	"recruit/authentication"
	"recruit/ent"
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

}

func initRoutes(client *ent.Client) {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "PUT", "GET", "OPTIONS"}
	config.MaxAge = 48 * time.Hour
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Cache-Control"}

	r.Use(cors.New(config))

	public := r.Group("/rect")
	r.Use(cors.New(config))
	r.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Recruitment APIs"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})

		//public := r.Group("/api")
	})
	admin := r.Group("/api/admin")
	admin.Use(validateToken)
	admin.Use(totopverify)

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
	//r.POST("/rect/examcalendar", CreateExamCalendar(client))
	r.GET("/rect/examcalendar/:id", GetExamCalendarID(client))
	r.POST("/rect/examcalendars/", CreateExamCalendar(client))
	//r.GET("/rect/examcalendars/Details/:id", getExamCalendarsWithDetails(client))

	r.GET("/rect/vacancyyears", getVacancyYears(client))
	r.POST("/rect/vacancyyears", CreateVacarncyYears(client))
	r.GET("/rect/vacancyyears/:id", GetVacancyYearID(client))
	//r.POST("/rect/vacancyyears/:id", UpdateVacancyYearID(client))
	r.GET("/rect/totp", topgenerate)

	r.GET("/rect/exampapers/:id", getExamPapers(client))
	r.GET("/rect/exampapers/ByExamCode/:id", GetExamPapersWithExamCode(client))
	r.GET("/rect/exampapers/All", getAllExamPapers(client))
	r.POST("/rect/exampapers/", InsertExamPapers(client))

	r.POST("/rect/employees/", CreateEmployee(client))
	r.GET("/rect/employees/All", getAllEmployees(client))
	r.PUT("/rect/employeeprofile/verify/:id", updateemployeeverifydetails(client))
	r.GET("rect/employees/:id", GetEmployeesByID(client))

	r.GET("/rect/employeecategories/", getAllEmployeeCategories(client))
	r.POST("/rect/employeecategories/", CreateEmployeeCategories(client))

	r.GET("/rect/employeedesignation/", getAllEmployeeDesignations(client))
	r.POST("/rect/employeedesignation/", CreateEmployeeDesignations(client))

	r.GET("/rect/employeedposts/", getAllEmployeePosts(client))
	r.GET("/rect/employeedposts/:id", GetEmpPostsByID(client))
	r.POST("/rect/employeedposts/", CreateEmployeePost(client))

	r.GET("/rect/employeeddisbilitytypes/", CreateDisabilityTypes(client))
	r.GET("/rect/employeeddisbilitytypes/:id", GetDisabilitysByID(client))
	r.POST("/rect/employeeddisbilitytypes/", CreateDisabilityTypes(client))

	r.POST("/rect/circles/", CreateCircleMasters(client))
	r.GET("/rect/circles/", getAllCircles(client))
	r.GET("/rect/circles/:id", GetCircleByID(client))

	r.POST("/rect/regions/", CreateRegionMasters(client))
	r.GET("/rect/regions/", getAllRegions(client))
	r.GET("/rect/regions/:id", GetRegionsByID(client))
	r.GET("/rect/regions/bycirclecode/:id", GetRegionsByCircleCode(client))

	r.POST("/rect/divisions/", CreateDivisionMasters(client))
	r.GET("/rect/divisions/", getAllDivisions(client))
	r.GET("/rect/divisions/:id", GetDivisionsByID(client))
	r.GET("/rect/divisions/byregioncode/:id", GetDivisionsByRegionCode(client))

	r.POST("/rect/facilities/", CreateFacilityMasters(client))
	r.GET("/rect/facilities/", getAllFacilities(client))
	r.GET("/rect/facilities/:id", GetFacilityByID(client))
	r.GET("/rect/facilities/bycirclecode/:id", GetFacilitiesByCircleCode(client))
	r.GET("/rect/facilities/bydivisioncode/:id", GetFacilitiesByDivisionCode(client))
	r.GET("/rect/facilities/byregioncode/:id", GetFacilitiesByRegionCode(client))

	r.GET("/totp", topgenerate)

	r.POST("/totopverify", totopverify)

	//r.GET("/token", GenerateToken)
	//r.GET("/verify", verifytoken)
	public.POST("/token", t)
	r.POST("/refreshtoken", validateRefreshToken)
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

func getAlDisabilityTypes(client *ent.Client) gin.HandlerFunc {
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

func GetDisabilitysByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		disabilityID, _ := strconv.ParseInt(id, 10, 32)

		disability, err := start.QueryEmployeePostsID(ctx, client, int32(disabilityID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": disability})

	}

	return gin.HandlerFunc(fn)
}
