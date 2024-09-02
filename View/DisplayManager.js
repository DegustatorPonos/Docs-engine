// Determines the parsing mode whe engine runs on
let Mode = 0;


async function DisplayContent(path) {
	console.log('path: ' + path)
	let returnStrings = [];
	await GetString("ReadFile?path=" + path	).then(result => {
		returnStrings = result.split('\n');
	}); 
	DetermineParsingMode(returnStrings.shift())
	console.log("Parsing mode: " + Mode);
	PushElemenetsOnDisplay(returnStrings);
}

// First returning string contains the info regarding parsing
function DetermineParsingMode(FirstString) {
	let _mode = FirstString.replace("mode ", "");
	Mode = _mode;
}

// Clears all elements on display and adds all given ones. Accepts list of elements
function PushElemenetsOnDisplay(elements) {
	let canvas = document.getElementById("canvas");
	ClearChildren(canvas);
	for(let i = 0; i < elements.length; i++) {
		if(elements[i] == "") continue;
		canvas.insertAdjacentHTML('beforeend', elements[i]);
		console.log(elements[i]);
	}
}

// Clears all elemets that are children of the given element
function ClearChildren(parentElement) {
	try {
	let els = parentElement.children;
	while(els.length != 0) {
		parentElement.removeChild(els[0]); // Thats fucked up honestly
	}
	} catch (ex) {
		console.log(ex)
	}
}