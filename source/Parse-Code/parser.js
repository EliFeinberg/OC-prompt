const fs = require("fs");
const csv = require("csv");

const result = {};
const keys = ["\"Text\"", "\"Description\""];

// Read data
const readStream = fs.createReadStream("whoami.txt");

// Parser
const parser = csv.parse({ delimiter: ":"});


parser.on("data", (chunk) => {

var text = chunk[0];
text = '"'+text+'"';
result[text] = {};
    for(let i = 1; i < chunk.length; i ++) {
        // var tempKey = JSON. stringify(keys[0]);
        result[text][keys[0]] = chunk[0];
        result[text][keys[1]] = chunk[1];
    }
});

parser.on("end", () => {
    console.log(result);
});

readStream.pipe(parser);


var myData = [];
rows.each(function (index) {
    var obj = { 
        id: $this.find('.elementOne').val(),
        name: $this.find('.elementTwo').text()
    };
    myData.push(obj);
});






// 1.run "curl -sL https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.0/install.sh -o install_nvm.sh"
// 2.run "bash install_nvm.sh"
// 3. copy end script and run
// 4. go to file and run "node parser.js"
// 5. turn all '"to "and turn all "' to "
// 6. turn all ' to "
// 7. turn all \" to '
// 8. turn all ': to '":