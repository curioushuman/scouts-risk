package fixtures

import "scouts-risk/internal/model"

var ControlsBefore = map[string][]model.Control{
"treeHouse": {
	{
		Label: "Cover the sharp edges",
		Id: "1",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Tree / branch inspection",
		Id: "2",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Review / check equipment prior to use",
		Id: "2",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Make sure safety rails/precautions exist",
		Id: "1",
		When: model.ControlWhenBefore,
	},
},
"bushy": {
	{
		Label: "Tree / branch inspection",
		Id: "2",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Bushfire inspection",
		Id: "2",
		When: model.ControlWhenBefore,
	},
},
"pack": {
	{
		Label: "All participants told to bring high boots, long pants",
		Id: "1",
		When: model.ControlWhenBefore,
	},
},
}