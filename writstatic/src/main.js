var content, nav;


function main()
{
	content = new Vue
	({
		el: "#content",
		data: {
			body: ""
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
		for(item in parsed.items)
		{
			item.text = "/r/" + item.text;
		}

		nav.data = parsed;
	}
	request.send();
}

function loadContent(name)
{
	
}

window.onload = main;