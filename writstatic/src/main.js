var content, nav;


function main()
{
	content = new Vue
	({
		el: "#content",
		data: {
			body: "",
			contentsEmpty: true
		}
	})

	nav = new Vue
	({
		el: "#nav",
		data: {
			items: []
		}
	})

	loadNav();
}

function loadNav()
{
	request = new XMLHttpRequest();
	request.open("GET", "/list");
	request.onload = function() 
	{
		var parsed = JSON.parse(request.response);
		var newItems = [];

		for(i = 0; i < parsed.items.length; i++)
		{
			const item = parsed.items[i];

			created = 
			{
				href: "/r/" + item,
				text: item
			};
			newItems.push(created);
		}

		nav.items = newItems;
	}
	request.send();
}

function loadContent(name)
{
	contentsEmpty = content.body == null || content.body.length <= 0;
}

window.onload = main;