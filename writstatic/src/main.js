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
			items: [],
			loading: false
		},
		methods: {
			navigate: function(event) {
				event.preventDefault();
				loadContent(event.target.attributes["href"].value);
			}
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
				href: "/f/" + item,
				text: item
			};
			newItems.push(created);
		}

		nav.items = newItems;
	}
	request.send();
}

function loadContent(path)
{
	nav.loading = true;

	request = new XMLHttpRequest();
	request.open("GET", path);
	request.onload = function()
	{
		content.body = request.response;
		content.contentsEmpty = content.body == null || content.body.length <= 0;
		nav.loading = false;
	}
	request.send();
}

window.onload = main;