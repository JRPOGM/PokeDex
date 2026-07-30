package pokeapi

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
	"github.com/JRPOGM/PokeDex/TheApi/pokecache"
)

const (
	baseURL = "http://pokeapi.co/api/v2"
)

type Client struct {
	cache		pokecache.Cache
	httpClient 	http.Client
}

type ResponseShallowLocations struct {
	Count 		int			`json:"count"`
	Next 		*string		`json:"next"`
	Previous 	*string		`json:"previous"`
	Results []struct {
		Name 	string		`json:"name"`
		URL 	string		`json:"url"`
	} `json:"results"`
}

type Location struct {
	EncounterMethodRates [] struct {
		EncounterMethod struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate		int `json:"rate"`
			Vaersion	struct {
				Name	string `json:"name"`
				URL		string `json:"url"`
			} `json:"version"`
		} `json:"version_details`
	} `json:"encounter_method_rates"`
	GameIndex		int `json:"game_index"`
	ID				int `json:"id"`
	Location struct {
		Name	string `json:"name"`
		URL		string `json:"url"`
	} `json:"location"`
	Name	string `json:"name"`
	Names []struct {
		Language struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"language"`
		Name	string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance				int `json:"chance"`
				ConditionValues		[]interface{} `json:"condition_values"`
				MaxLevel			int `json:"max_level"`
				Method struct {
					Name	string `json:"name"`
					URL		string `json:"url"`
				} `json:"method"`
				MinLevel			int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance	int `json:"max_chance"`
			Version struct {
				Name	string `json:"name"`
				URL		string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"ability"`
		IsHidden	bool `json:"is_hidden"`
		Slot		int `json:"slot"`
	} `json:"abilities"`
	BaseExperience	int `json:"base_experience"`
	Forms []struct {
		Name	string `json:"name"`
		URL		string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex	int `json:"game_index"`
		Version struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height						int `json:"height"`
	HeldItems					[]interface{} `json:"held_items"`
	ID							int `json:"id"`
	IsDefault					bool `json:"is_default"`
	LocationAreaEncounters		string `json:"location_area_encounters"`
	Moves []struct {
		Move struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt	int `json:"level_learned_at"`
			MoveLearnedAt struct {
				Name	string `json:"name"`
				URL		string `json:"url"`
			} `json:"move_learned_at"`
			VersionGroup struct {
				Name	string `json:"name"`
				URL		string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name		string `json:"name"`
	Order		int `json:"order"`
	PastTypes	[]interface{} `json:"past_types"`
	Species struct {
		Name	string `json:"name"`
		URL		string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault			string `json:"back_default"`
		BackFemale			interface{} `json:"back_female"`
		BackShiny			string `json:"back_shiny"`
		BackShinyFemale		interface{} `json:"back_shiny_female"`
		FrontDefault		string `json:"front_default"`
		FrontFemale			interface{} `json:"front_female"`
		FrontShiny			string `json:"front_shiny"`
		FrontShinyFemale	interface{} `json:"front_shiny_female"`
		Other struct {
			DeramWorld struct {
				FrontDefault	string `json:"front_default"`
				FrontFemale		interface{} `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault		string `json:"front_default"`
				FrontFemale			interface{} `json:"front_female"`
				FrontShiny			string `json:"front_shiny"`
				FrontShintFemale	interface{} `json:front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault	string `json:"front_default"`
				FrontShiny		string `json:"front_shiny"`
			} `json:"official_artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault			string `json:"back_default"`
					BackGray			string `json:"back_gray"`
					BackTransparent		string `json:"back_transparent"`
					FrontDefault		string `json:"front_default"`
					FrontGray			string `json:"front_gray"`
					FrontTransparent	string `json:"front_transparent"`
				} `json:"red_blue"`
				Yellow struct {
					BackDefault			string `json:"back_default"`
					BackGray			string `json:"back_gray"`
					BackTransparent		string `json:"back_transparent"`
					FrontDefault		string `json:"front_default"`
					FrontGray			string `json:"front_gray"`
					FrontTransparent	string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation_i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault				string `json:"back_default"`
					BackShiny				string `json:"back_shiny"`
					BackShinyTransparent	string `json:"back_shinytransparent"`
					BackTransparent			string `json:"back_transparent"`
					FrontDefault			string `json:"front_default"`
					FrontShiny				string `json:"front_shiny"`
					FrontShinyTransparent	string `json:"front_shiny_transparent"`
					FrontTransparent		string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault				string `json:"back_default"`
					BackShiny				string `json:"back_shiny"`
					FrontDefault			string `json:"front_default"`
					FrontShiny				string `json:"front_shiny"`
					FrontTransparent		string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault				string `json:"back_default"`
					BackShiny				string `json:"back_shiny"`
					FrontDefault			string `json:"front_default"`
					FrontShiny				string `json:"front_shiny"`
					FrontTransparent		string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation_ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault	string `json:"front_default"`
					FrontShiny		string `json:"front_shiny"`
				} `json:"emerald"`
				RubySapphire struct {
					BackDefault				string `json:"back_default"`
					BackShiny				string `json:"back_shiny"`
					FrontDefault			string `json:"front_default"`
					FrontShiny				string `json:"front_shiny"`
				} `json:"ruby_sapphire"`
				FireredLeafgreen struct {
					BackDefault				string `json:"back_default"`
					BackShiny				string `json:"back_shiny"`
					FrontDefault			string `json:"front_default"`
					FrontShiny				string `json:"front_shiny"`
				} `json:firered_leafgreen"`
			} `json:"generation_iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault			string `json:"back_default"`
					BackFemale			interface{} `json:"back_female"`
					BackShiny			string `json:"back_shiny"`
					BackShinyFemale		interface{} `json:"back_shiny_female"`
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"diamond_pearl"`
				Platinum struct {
					BackDefault			string `json:"back_default"`
					BackFemale			interface{} `json:"back_female"`
					BackShiny			string `json:"back_shiny"`
					BackShinyFemale		interface{} `json:"back_shiny_female"`
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"platinum"`
				HeartgoldSoulsilver struct {
					BackDefault			string `json:"back_default"`
					BackFemale			interface{} `json:"back_female"`
					BackShiny			string `json:"back_shiny"`
					BackShinyFemale		interface{} `json:"back_shiny_female"`
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"heartgold_soulsilver"`
			} `json:"generation_iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault			string `json:"back_default"`
						BackFemale			interface{} `json:"back_female"`
						BackShiny			string `json:"back_shiny"`
						BackShinyFemale		interface{} `json:"back_shiny_female"`
						FrontDefault		string `json:"front_default"`
						FrontFemale			interface{} `json:"front_female"`
						FrontShiny			string `json:"front_shiny"`
						FrontShinyFemale	interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault			string `json:"back_default"`
					BackFemale			interface{} `json:"back_female"`
					BackShiny			string `json:"back_shiny"`
					BackShinyFemale		interface{} `json:"back_shiny_female"`
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"black_white"`
				BlacktwoWhitetwo struct {
					Animated struct {
						BackDefault			string `json:"back_default"`
						BackFemale			interface{} `json:"back_female"`
						BackShiny			string `json:"back_shiny"`
						BackShinyFemale		interface{} `json:"back_shiny_female"`
						FrontDefault		string `json:"front_default"`
						FrontFemale			interface{} `json:"front_female"`
						FrontShiny			string `json:"front_shiny"`
						FrontShinyFemale	interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault			string `json:"back_default"`
					BackFemale			interface{} `json:"back_female"`
					BackShiny			string `json:"back_shiny"`
					BackShinyFemale		interface{} `json:"back_shiny_female"`
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"blacktwo_whitetwo"`
			} `json:"generation_v"`
			GenerationVi struct {
				XY struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"x_y"`
				OmegarubyAlphaSapphire struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"omegaruby_alphasapphire"`
			} `json:"generation_vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
				} `json:"icons"`
				SunMoon struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"sun_moon"`
				UltrasunUltramoon struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"ultrasun_ultramoon"`
			} `json:"generation_vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
				} `json:"icons"`
				SwordShield struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"sword_shield"`
				BrilliantdiamondShiningpearl struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"brilliantdiamond_shiningpearl"`
				LegendsArceus struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"legends_arceus"`
			} `json:"generation_viii"`
			GenerationIx struct {
				Icons struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
				} `json:"icons"`
				ScarletViolet struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"scarlet_violet"`
				LegendsZA struct {
					FrontDefault		string `json:"front_default"`
					FrontFemale			interface{} `json:"front_female"`
					FrontShiny			string `json:"front_shiny"`
					FrontShinyFemale	interface{} `json:"front_shiny_female"`
				} `json:"legends_z_a"`
			} `json:"generation_ix"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat	int `json:"base_stat"`
		Effort		int `json:"effort"`
		Stat struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot		int `json:"slot"`
		Type struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight	int `json:"weight"`
}

func NewClient(cacheInterval, timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ListLocations(pageURL *string) (ResponseShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if value, ok := c.cache.Get(url); ok {
		localResponse := ResponseShallowLocations{}
		err := json.Unmarshal(value, &localResponse)
		if err != nil {
			return ResponseShallowLocations{}, err
		}
		return localResponse, nil
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	localResponse := ResponseShallowLocations{}
	err = json.Unmarshal(data, &localResponse)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	c.cache.Add(url, data)
	return localResponse, nil
}

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	if value, ok := c.cache.Get(url); ok {
		localResponse := Location{}
		err := json.Unmarshal(value, &localResponse)
		if err != nil {
			return Location{}, err
		}
		return localResponse, nil
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}
	localResponse := Location{}
	err = json.Unmarshal(data, &localResponse)
	if err != nil {
		return Location{}, err
	}
	c.cache.Add(url, data)
	return localResponse, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if value, ok := c.cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(value, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonResponse := Pokemon{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	return pokemonResponse, nil
}