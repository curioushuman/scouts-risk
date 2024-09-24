package fixtures

import "scouts-risk/internal/model"

var Locations = map[string]model.Location{
"treeHouse": {
	Label: "Mac's Tree House",
  Id: "treeHouse",
	LocType: model.LocationTypeEquipment,
	Hazards: []model.Hazard{
		{
			Label: "Exposed sharp edges / points",
			Id: "1",
			Controls: []model.Control{
				{
					Label: "Cover the sharp edges",
					Id: "11",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "12",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "13",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Falling branches",
			Id: "2",
			Controls: []model.Control{
				{
					Label: "Tree / branch inspection",
					Id: "22",
					When: model.ControlWhenBefore,
				},
				{
					Label: "First aid",
					Id: "23",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Loss of structural integrity",
			Id: "3",
			Controls: []model.Control{
				{
					Label: "Ensure installation by skilled person",
					Id: "31",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Review / check equipment prior to use",
					Id: "32",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Review / maintain equipment regularly",
					Id: "33",
					When: model.ControlWhenRegular,
				},
				{
					Label: "First aid",
					Id: "34",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "35",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Falling from height",
			Id: "4",
			Controls: []model.Control{
				{
					Label: "Make sure safety rails/precautions exist",
					Id: "41",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Utilise necessary safety equipment; for height",
					Id: "42",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "43",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "44",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "45",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Equipment failure",
			Id: "6",
			Controls: []model.Control{
				{
					Label: "Ensure installation by skilled person",
					Id: "61",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Review / check equipment prior to use",
					Id: "62",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Review / maintain equipment regularly",
					Id: "63",
					When: model.ControlWhenRegular,
				},
				{
					Label: "First aid",
					Id: "64",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "65",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Improper use",
			Id: "7",
			Controls: []model.Control{
				{
					Label: "Sufficient safety briefing",
					Id: "71",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Monitor scouts while in use",
					Id: "72",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "73",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "74",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Incorrect configuration",
			Id: "8",
			Controls: []model.Control{
				{
					Label: "Ensure installation by skilled person",
					Id: "81",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Review / check equipment prior to use",
					Id: "82",
					When: model.ControlWhenBefore,
				},
				{
					Label: "First aid",
					Id: "83",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "84",
					When: model.ControlWhenResponse,
				},
			},
		},
	},
},
"bushy": {
	Label: "Bushy / forested area",
  Id: "bushy",
	LocType: model.LocationTypeArea,
	Hazards: []model.Hazard{
		{
			Label: "Snakes",
			Id: "9",
			Controls: []model.Control{
				{
					Label: "All participants told to bring high boots, long pants",
					Id: "91",
					When: model.ControlWhenBefore,
				},
				{
					Label: "All participants to wear high boots, long pants in bushy area",
					Id: "92",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Stay with/behind a leader where possible",
					Id: "93",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Check the area prior to use",
					Id: "94",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "95",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "96",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "97",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Spiders",
			Id: "10",
			Controls: []model.Control{
				{
					Label: "Check the area prior to use",
					Id: "101",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "102",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "103",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "104",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Leeches",
			Id: "11",
			Controls: []model.Control{
				{
					Label: "All participants told to bring high boots, long pants",
					Id: "111",
					When: model.ControlWhenBefore,
				},
				{
					Label: "All participants to wear high boots, long pants in bushy area",
					Id: "112",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "113",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Check for ticks & leeches after use",
					Id: "114",
					When: model.ControlWhenAfter,
				},
				{
					Label: "First aid",
					Id: "115",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Other fauna",
			Id: "12",
			Controls: []model.Control{
				{
					Label: "Check the area prior to use",
					Id: "122",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "123",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "124",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "125",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Insects",
			Id: "13",
			Controls: []model.Control{
				{
					Label: "Check the area prior to use",
					Id: "131",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "132",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "133",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "134",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Ticks",
			Id: "14",
			Controls: []model.Control{
				{
					Label: "All participants told to bring high boots, long pants",
					Id: "141",
					When: model.ControlWhenBefore,
				},
				{
					Label: "All participants to wear high boots, long pants in bushy area",
					Id: "142",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "143",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Check for ticks & leeches after use",
					Id: "144",
					When: model.ControlWhenAfter,
				},
				{
					Label: "First aid",
					Id: "145",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "146",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Outdoor trip hazards",
			Id: "15",
			Controls: []model.Control{
				{
					Label: "Remove or mark trip hazards",
					Id: "151",
					When: model.ControlWhenBefore,
				},
				{
					Label: "First aid",
					Id: "152",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Falling trees",
			Id: "16",
			Controls: []model.Control{
				{
					Label: "Tree / branch inspection",
					Id: "161",
					When: model.ControlWhenBefore,
				},
				{
					Label: "First aid",
					Id: "162",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "163",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Falling branches",
			Id: "17",
			Controls: []model.Control{
				{
					Label: "Tree / branch inspection",
					Id: "171",
					When: model.ControlWhenBefore,
				},
				{
					Label: "First aid",
					Id: "172",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "173",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Bushfire",
			Id: "18",
			Controls: []model.Control{
				{
					Label: "Bushfire inspection",
					Id: "181",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Re-assess / cancel based on weather",
					Id: "182",
					When: model.ControlWhenBefore,
				},
			},
		},
	},
},
"macsPlace": {
	Label: "Mac's place",
  Id: "macsPlace",
	LocType: model.LocationTypeLocation,
	Hazards: []model.Hazard{
		{
			Label: "Outdoor trip hazards",
			Id: "19",
			Controls: []model.Control{
				{
					Label: "Remove or mark trip hazards",
					Id: "191",
					When: model.ControlWhenBefore,
				},
				{
					Label: "First aid",
					Id: "192",
					When: model.ControlWhenResponse,
				},
			},
		},
	},
},
}

var LocationArea = model.Location{
	Label: "Pool / dam / body of water",
  Id: "pool",
	LocType: model.LocationTypeArea,
	Hazards: []model.Hazard{
		{
			Label: "Falling into water",
			Id: "pool1",
			Controls: []model.Control{
				{
					Label: "Fence off area",
					Id: "pool11",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Sufficient safety briefing",
					Id: "pool12",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "pool12",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "pool24",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Improper use",
			Id: "pool2",
			Controls: []model.Control{
				{
					Label: "Sufficient safety briefing",
					Id: "pool21",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Monitor scouts while in use",
					Id: "pool22",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "pool23",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "pool24",
					When: model.ControlWhenResponse,
				},
			},
		},
	},
}