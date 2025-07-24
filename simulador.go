package main

import (
	//"context"
	"fmt"
	"time"
	"math/rand"
	"math"
  "context"
	"sync"
	"flag"
	"github.com/influxdata/influxdb-client-go/v2"
  "github.com/influxdata/influxdb-client-go/v2/api/write"

)

type Car struct {
	ID 			string
	RPM 		int 
	Speed 	float64
	TPS 		float64
	Temp 		float64
	Lambda 	float64
	KM 			float64
	Mu 			sync.Mutex
}
func (c * Car)RunSimulation(){
	for {
	c.Mu.Lock()
		c.TPS = math.Max(0, math.Min(1, c.TPS+(rand.Float64()-0.5)*0.2))

		// Ajustar RPM según TPS
		if c.TPS > 0.1 {
			c.RPM += int(rand.Float64()*1000*c.TPS) - 300
		} else {
			c.RPM -= 250 // baja por inercia
		}

		// Limitar RPM
		if c.RPM < 800 {
			c.RPM = 800
		}
		if c.RPM > 7000 {
			c.RPM = 7000
		}

		// Ajustar velocidad (simulando inercia y frenado)
		if c.TPS > 0.05 {
			c.Speed += float64(c.RPM) / 5000
		} else {
			c.Speed -= 2.5 // frenado
		}

		// Limitar velocidad
		if c.Speed < 0 {
			c.Speed = 0
		}
		if c.Speed > 180 {
			c.Speed = 180
		}

		// Temperatura del motor
		if c.RPM > 3000 && c.TPS > 0.4 {
			c.Temp += 0.3
		} else if c.Temp > 75 {
			c.Temp -= 0.1
		}

		if c.Temp < 60 {
			c.Temp = 60
		}
		if c.Temp > 110 {
			c.Temp = 110
		}

		// Simular lambda (mezcla aire/combustible)
		c.Lambda = 0.85 + rand.Float64()*0.3 // entre 0.85 y 1.15

		// Simular odómetro
		c.KM += c.Speed / 3600 // km por segundo
	c.Mu.Unlock()
}
}

func main (){
    carID := flag.String("id", "1", "ID del vehículo")
    flag.Parse()
    c := Car{
        ID:        *carID,
        RPM:       800.0,
        Speed: 0.0,
        TPS:       0.0,
        Temp: 75.0,
        Lambda:    1.0,
        KM: 12345.67,
    }

	url := "http://localhost:8086"
	token :=	"fcB4y7pynbxa3tnn-laLye0H-LeGSJ_uehM__H2VWskV0_orCoJeYNjGSsiADjSp8GjyvlkM8V69PKSnlFhuxg=="
	org := "mi-org"
	bucket := "telemetry"

	client := influxdb2.NewClient(url, token)
  writeAPI := client.WriteAPIBlocking(org, bucket)
	defer client.Close()

	ticker := time.NewTicker(1 * time.Second)
	go c.RunSimulation()

  for range ticker.C{ 
		fmt.Println(c)
        point := write.NewPointWithMeasurement("car_data").
            AddTag("car_id", c.ID).
            AddField("rpm", c.RPM).
            AddField("speed", c.Speed).
            AddField("tps", c.TPS).
            AddField("Engine_Temp", c.Temp).
            AddField("lambda", c.Lambda).
            AddField("TotalKM", c.KM).
            SetTime(time.Now())
		    err := writeAPI.WritePoint(context.Background(), point)

 if err != nil {
            fmt.Println("Error al escribir en InfluxDB:", err)
        } else {
            fmt.Printf(">> [%s] rpm: %d | vel: %.1f km/h | tps: %.2f\n", c.ID, c.RPM, c.Speed, c.TPS)
        }

	}

					
}
