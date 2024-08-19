package repositories

import (
    "testing"
    "time"

    domain "github.com/legend123213/go_togo/Task08/task-manager/Domain"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_Gettask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        Newrepouser := NewTaskService(userCollection)
        expectedUser := &domain.Task{
            Title:       "Test Task",
            Description: "Test Description",
            Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
            Status:      "Pending",
        }

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
            {"_id", expectedUser.ID},
            {"title", expectedUser.Title},
            {"description", expectedUser.Description},
            {"due_date", expectedUser.Due_date},
            {"status", expectedUser.Status},
        }))
        userResponse, err := Newrepouser.SGetTask(expectedUser.ID.Hex(), "")
        assert.Nil(t, err)
        assert.Equal(t, *expectedUser, *userResponse)
    })
}

func Test_Addtask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        Newrepotask := NewTaskService(userCollection)
        id := primitive.NewObjectID()
        newtask := &domain.Task{
            Title:       "Test Task",
            Description: "Test Description",
            Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
            Status:      "Pending",
            UserID:      id,
        }

        mt.AddMockResponses(mtest.CreateSuccessResponse())

        insertedUser, err := Newrepotask.SAddTask(newtask)
        assert.Nil(t, err)
        assert.Equal(t, domain.Task{
            ID:          insertedUser.ID,
            Title:       "Test Task",
            Description: "Test Description",
            Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
            Status:      "Pending",
            UserID:      id,
        }, *insertedUser)
    })
}

func Test_updateTask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        Newrepotask := NewTaskService(userCollection)
        id := primitive.NewObjectID()
        newtask := &domain.Task{
            ID:          id,
            Title:       "Test Task",
            Description: "Test Description",
            Due_date:    time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
            Status:      "Pending",
            UserID:      id,
        }
        mt.AddMockResponses(bson.D{
            {"ok", 1},
            {"value", bson.D{
                {"_id", newtask.ID},
                {"title", newtask.Title},
                {"description", newtask.Description},
                {"due_date", newtask.Due_date},
                {"status", newtask.Status},
            }},
        })

        updatedUser, err := Newrepotask.SEditTask(primitive.NewObjectID().Hex(), newtask)

        assert.Nil(t, err)
        assert.Equal(t, *newtask, *updatedUser)
    })
}

func Test_DeleteTask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("success", func(mt *mtest.T) {
        userCollection := mt.Coll
        Newrepotask := NewTaskService(userCollection)
        id := primitive.NewObjectID()

        mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})

        err := Newrepotask.SDeleteTask(id.Hex())
        assert.Nil(t, err)
    })
}


func Test_GetAllTasks(t *testing.T) {
        mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

        mt.Run("success", func(mt *mtest.T) {
            userCollection := mt.Coll
            Newrepotask := NewTaskService(userCollection)

            expectedTasks := []*domain.Task{
                {
                    ID:          primitive.NewObjectID(),
                    Title:       "Task 1",
                    Description: "Description 1",
                    Due_date:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
                    Status:      "Pending",
                    UserID:      primitive.NewObjectID(),
                },
                {
                    ID:          primitive.NewObjectID(),
                    Title:       "Task 2",
                    Description: "Description 2",
                    Due_date:    time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
                    Status:      "Completed",
                    UserID:      primitive.NewObjectID(),
                },
            }

            mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch, bson.D{
                {"_id", expectedTasks[0].ID},
                {"title", expectedTasks[0].Title},
                {"description", expectedTasks[0].Description},
                {"due_date", expectedTasks[0].Due_date},
                {"status", expectedTasks[0].Status},
                {"user_id", expectedTasks[0].UserID},
            }, bson.D{
                {"_id", expectedTasks[1].ID},
                {"title", expectedTasks[1].Title},
                {"description", expectedTasks[1].Description},
                {"due_date", expectedTasks[1].Due_date},
                {"status", expectedTasks[1].Status},
                {"user_id", expectedTasks[1].UserID},
            }))

            tasks, err := Newrepotask.SGetTasks("")
            assert.Nil(t, err)

            assert.Equal(t, expectedTasks[0].ID, tasks[0].ID)
            assert.Equal(t, expectedTasks[0].Title, tasks[0].Title)
            assert.Equal(t, expectedTasks[0].Description, tasks[0].Description)
            assert.Equal(t, expectedTasks[0].Due_date, tasks[0].Due_date)
            assert.Equal(t, expectedTasks[0].Status, tasks[0].Status)
            assert.Equal(t, expectedTasks[0].UserID, tasks[0].UserID)

            assert.Equal(t, expectedTasks[1].ID, tasks[1].ID)
            assert.Equal(t, expectedTasks[1].Title, tasks[1].Title)
            assert.Equal(t, expectedTasks[1].Description, tasks[1].Description)
            assert.Equal(t, expectedTasks[1].Due_date, tasks[1].Due_date)
            assert.Equal(t, expectedTasks[1].Status, tasks[1].Status)
            assert.Equal(t, expectedTasks[1].UserID, tasks[1].UserID)
        })
    }

