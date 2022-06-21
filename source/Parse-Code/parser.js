const fs = require("fs");
const csv = require("csv");

const result = {};
const keys = ["\"Text\"", "\"Description\""];

// Read data
const readStream = fs.createReadStream("GlobalOp.txt");

// Parser
const parser = csv.parse({ delimiter: ":" });


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

// const json = JSON.stringify(result);
// fs.writeFile('test.json', json, 'utf8')

readStream.pipe(parser);

// 1.run "curl -sL https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.0/install.sh -o install_nvm.sh"
// 2.run "bash install_nvm.sh"
// 3. copy end script and run
// 4. go to file and run "node parser.js"
// 5. turn all '" to " and turn all "' to "
// 6. turn all ' to "