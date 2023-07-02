package handlers

import (
	"context"
	"net/http"
	"recruit/ent"
	start "recruit/start"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*func SomeHandler(db *sql.DB) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        // Your handler code goes in here - e.g.
        rows, err := db.Query(...)

        c.String(200, results)
    }

    return gin.HandlerFunc(fn)
}*/

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

/*func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}

	return "Unknown error"
}*/

func CreateExam(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newExam := new(ent.Exam)

		if err := gctx.ShouldBindJSON(&newExam); err != nil {
			/*var ve validator.ValidationErrors

			if errors.As(err, &ve) {
				out := make([]ErrorMsg, len(ve))
				for i, fe := range ve {
					out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
				}
				gctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			}*/

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

		newExam, err := start.CreateExam(client, newExam)
		//fmt.Println("newExam", newExam)
		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam")

	}
	return gin.HandlerFunc(fn)
}

func GetExamID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		examID, _ := strconv.ParseInt(id, 10, 32)

		exam, err := start.QueryExamID(ctx, client, int32(examID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exam})

	}

	return gin.HandlerFunc(fn)
}
