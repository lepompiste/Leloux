;(
function(){
	function createElement(type, props, ...children) {
		return {
			type: type,
			props: {
				...props,
				children: children.map(child => {
					return typeof child == "object" ? child : createTextElement(child)
				})
			}
		}
	}

	function createTextElement(text) {
		return {
			type: "TEXT_ELEMENT",
			props: {
				nodeValue: text,
				children: []
			}
		}
	}

	function render(container, element) {
		const vdom = element.type == "TEXT_ELEMENT" ? document.createTextNode(element.props.nodeValue) : document.createElement(element.type)
		
		Object.keys(element.props).forEach(name => {
			if (name != "children") {
				vdom[name] = element.props[name]
			}
		})

		element.props.children.forEach(child => {
			render(vdom, child)
		});

		container.appendChild(vdom)
	}

	Leloux = createElement
	Leloux.createElement = createElement
	Leloux.createTextElement = createTextElement
	Leloux.render = render

	window.Leloux = Leloux
	window.l = Leloux
})()