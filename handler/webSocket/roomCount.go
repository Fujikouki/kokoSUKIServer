package webSocket

import (
	"encoding/json"
	"net/http"
)

type RoomCount struct {
	Count    int
	RoomList []string
	RoomSize []int
}

func roomCount(w http.ResponseWriter, r *http.Request) {

	count := len(rooms)
	roomList := make([]string, 0, count)
	roomSize := make([]int, 0, count)
	for k := range rooms {
		roomList = append(roomList, k)
		roomSize = append(roomSize, len(rooms[k].clients))
	}

	roomCount := &RoomCount{
		Count:    count,
		RoomList: roomList,
		RoomSize: roomSize,
	}

	err := json.NewEncoder(w).Encode(roomCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
