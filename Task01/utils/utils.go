package utils


func Average(subjects map[string]float64) float64{
	totol_score := float64(0)
   for _,val := range subjects{
		totol_score+=val
	}
	return totol_score/float64(len(subjects))
}