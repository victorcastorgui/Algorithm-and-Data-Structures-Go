package main

const maximumUnrecommendedCriminals int = 5
const maximumPassengerInVehicle int = 3

type Person struct {
	Name          string
	CriminalScore int
}

func LastDayInJail(criminals []Person, chosenPerson string) (onTransport []Person, waiting []Person) {
	var releasedCriminals []Person
	var vehicle []Person
	var wait []Person

	if len(criminals) == 0 {
		return nil, nil
	}

	sort(criminals)
	for i, val := range criminals {
		if i < maximumUnrecommendedCriminals {
			releasedCriminals = append(releasedCriminals, val)
		} else {
			break
		}
	}

	if chosenPerson != "" {
		for _, val := range criminals {
			if val.Name == chosenPerson {
				releasedCriminals = append(releasedCriminals, val)
			}
		}
	}

	for i, val := range releasedCriminals {
		if i < maximumPassengerInVehicle {
			vehicle = append(vehicle, val)
		} else {
			wait = append(wait, val)
		}
	}

	return vehicle, wait
}

func sort(criminals []Person) {
	for i := 0; i < len(criminals)-1; i++ {
		for j := 0; j < len(criminals)-i-1; j++ {
			if criminals[j].CriminalScore > criminals[j+1].CriminalScore {
				criminals[j], criminals[j+1] = criminals[j+1], criminals[j]
			} else if criminals[j].CriminalScore == criminals[j+1].CriminalScore {
				if criminals[j].Name > criminals[j+1].Name {
					criminals[j], criminals[j+1] = criminals[j+1], criminals[j]
				}
			}
		}
	}
}
