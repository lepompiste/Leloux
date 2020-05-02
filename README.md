# Leloux
basic component, routes and request library

## Useful informations

`l` is a shortcut to `Leloux` function.

`Leloux` is a shortcut to `Leloux.createElement` method.



## Elements

You can create elements like this :

```javascript
l("h1", {
    "id": "my_id",
    "class": "class1 class2",
    "events": {
        "click": function(e) {
            console.log("clicked")
        }
    }
}, "MY TITLE")
// this code creates an h1 element, with text "MY TITLE", id "my_id", class  "class1" and "class2", and a click events who execute the function
```

You can even create trees :

```javascript
l("div", {},
	l("h1", {}, "TITLE"),
	l("p", {}, "PARAGRAPH TEXT")
)
```



## Components

A component is just an object, with a `view` method.

An `init` method can also be provided.

```javascript
var Component = {
    init: function () {
        
    },
    view: function () {
        return l()
    }
}
```



## Rendering

You can render elements by using the method `l.render(MOUNTPOINT, COMPONENT)`, with MOUNTPOINT as a dom element, and COMPONENT a Leloux component.

Example :

```javascript
l.render(document.getElementById("view"), MyComponent)
```

You can also `mount` components, which is basically the same as `render`, but all mounted components are redrew when `l.redrawAll` method is called.

## Routes

Leloux includes a basic router. You can use it like this:

```javascript
l.routes.def(MOUNTPOINT, defaultRoute, {
    "/route1": Component1,
    "/route2": Component2
})
```

The route is called with the URL:

```
http://www.mydomain.com/file.html#!/route1
```



You can provide a parameter to the URL, after the route after `!`, and can be got using `l.routes.getParam` method.

Example :

```
http://www.mydomain.com/file.html#!/route1!parameter
```

Component is reloaded every time the route or parameter change.



Here are different useful methods :

```javascript
l.routes.onchangeroute = function() {} // add a callback on changing route
l.routes.reload() // reload the current route, reloading the component
l.routes.goto(ROUTE, PARAM) // go to the specified route, using the optional PARAM
```



By default, components are mounted to the MOUNTPOINT, but you can make it just rendered to preserve performances, by changing `l.routes.mount` to `false`.

You can also change the symbols of routes (default "#!") and parameter (default "!").

All those parameters have to be changed (or not) before the call to `l.routes.def`



## Requests

Leloux provides a basic requests tool, that you can use like that :

```javascript
l.requests.make(METHOD, URL, PROPS, ISJSON) // returns a Promise
```

`METHOD` is "GET", "POST", etc...

`URL` is... the URL

`PROPS` is an object :

​	`PROPS.query` are query parameters (key : value)

​	`PROPS.body` is the body of the request (key : value).

​	`PROPS.bodyRaw` is sent if and only if `PROPS.body` is undefined. It is send as is.

​	`PROPS.events` add events to the xhr request ("load", "error", ...).

​	`PROPS.uploadEvents` add events while upload to the xhr requests ("progress",  ...).

​	`PROPS.headers` are the headers of the request

​	`PROPS.useFormData` define if FormData is used to parse data (multipart/form-data encoding like) or not (application/www-x-form-urlencoded encoding like).

`ISJSON` is a Boolean that defines whether the response should be parsed as JSON or not.



You can also use `l.requests.makej` that calls `l.requests.make` with ISJSON = `true`



Example :

```javascript
l.requests.makej("GET", "./api/test", {
    "query": {
        "token": token
    }
}).then(function (jsonData) {
    console.log(jsonData)
})
```



Example using route and requests :

```javascript
var Component = {
    init: () => {
        m.requests.makej("GET", "/api/getText", {
            token: token
        })
        .then(data => {
            l.renderElement(document.getElementById("test_section"), l("p", {}, data.r))
        })
    },
    view: () => {
        return [
            l("h1", {}, "title"),
            l("section", {id: "test_section"}, "")
        ]
    }
}

l.routes.def(MOUNTPOINT, "/", {
    "/": Component
})
```



Example upload multiple files :

```html
<form action="./upload" enctype="multipart/form-data" method="POST" id="uploader">
	<input type="file" name="uploads" multiple id="uploads">
	<input type="submit" value="Submit">
</form>
	
<script type="text/javascript" defer>
	document.querySelector("#uploader").addEventListener("submit", function (evt) {
		evt.preventDefault()

		filesData = new FormData()
		for (i = 0; i < document.querySelector("#uploads").files.length; i++) {
            // Add all files to the form data under the same name
			filesData.append("uploads", document.querySelector("#uploads").files[i])
		}

		l.requests.makej("POST", "./upload", {
			headers: { // Content type to null to override default one. Maybe can be set as multipart/form-data but it works like this
				"Content-type": null
			},
			uploadEvents: {
				"progress": (event) => { // display progress in the console
					console.log(event.loaded + "/" + event.total)
				} 
			},
			bodyRaw: filesData
		}).then((data) => {
			console.log(data)
		})
	})
</script>
```
