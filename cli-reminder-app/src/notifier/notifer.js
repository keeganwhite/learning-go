const express = require("express");
const bodyParser = require("body-parser");
const notifier = require("node-notifier")
const {response} = require("express");
const port = process.env.PORT || 9000;

const app = express();
app.use(bodyParser.json());

app.get("/health", (req, resp) => resp.status(200).send());
app.post("/notify", (req, resp) => {
    notify(req.body, reply => resp.send(reply))
});
app.listen(port, () => console.log(`Server has started and is running on port: ${port}`));

const notify = ({title, message}, callback) => {
    notifier.notify(
        {
            title: title || "Unknown title",
            message: message || "Unknown message",
            sound: true,
            wait: true,
            reply: true,
            closeLabel: "Completed",
            timeout: 15
        },
        (error, response, reply) => {
            callback(reply)
        }
    )
};