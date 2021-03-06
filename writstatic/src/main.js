var content = null;
var nav = null;

function main()
{
	content = new Vue
	({
		el: "#content",
		data: {
			body: "",
			contentsEmpty: true,
			selected: ""
		}
	})

	try
	{
		nav = new Vue
		({
			el: "#nav",
			data: {
				items: [],
				loading: false,
				selected: ""
			},
			methods: {
				navigate: function(event) {
					event.preventDefault();

					var path = event.target.attributes["text"].value;
					loadContent(path, true);
				}
			}
		})

		loadNav(checkCurrentURLContent);	
	}
	catch(err){}
}

function loadNav(callback)
{
	request = new XMLHttpRequest();
	request.open("GET", "/list", true);
	request.onload = function() 
	{
		var parsed = JSON.parse(request.response);
		var newItems = [];

		for(i = 0; i < parsed.items.length; i++)
		{
			const item = parsed.items[i];
			if(item == "_about")
				continue;

			created = {text: item};
			newItems.push(created);
		}

		nav.items = newItems;

		// for unknown reasons, we have to chain (rather than in parallel) xhr requests.
		if(callback)
			callback();
	}
	request.send();
}

function loadContent(title, addToHistory)
{
	request = new XMLHttpRequest();
	request.open("GET", "/f/" + title, true);
	request.onload = function()
	{
		content.body = request.response;
		content.contentsEmpty = content.body == null || content.body.length <= 0;
		
		if(!content.contentsEmpty)
		{
			document.title = title;
			window.scrollTo(0, 0);
		}

		if(nav != null)
		{
			nav.loading = false;
			nav.selected = title;
		}
	}
	request.send();

	if(nav != null)
		nav.loading = true;

	if(addToHistory)
		history.pushState(null, title, "/a/" + title);
}

function checkCurrentURLContent()
{
	// check to see if we're loading into a guide directly on load.
	if(window.location.pathname.startsWith("/a/"))
	{
		var title = window.location.pathname.replace("/a/", "");
		title = decodeURIComponent(title);
		loadContent(title, false);
		return;
	}

	if(window.location.pathname.startsWith("/about"))
	{
		loadContent("_about", false);
		return;
	}
}

window.onload = main;
window.onpopstate = checkCurrentURLContent;