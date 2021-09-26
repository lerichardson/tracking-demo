# Demo of ad tracking software
I have a horrible fascination with the adtech machine. I've always wanted to see how it worked, so I did it myself.

This is purely educational. I have no intents of actually tracking people, and while the license does allow for it, I strongly discourage actually using it for that purpose.

To learn more, [follow my Medium](https://medium.litbelb.com), where I will write an in-depth article about why and how I made this.

## Contributing
PRs are welcome  
Assuming you already have go installed, run
```sh
sudo go run main.go
```
The standard port is 3000. You can change it by changing `":3000"` in the `app.Listen` function in `main.go` to `":[YOUR_PORT_HERE]"`. The demo is available at [baseurl/demo](http://localhost:3000/demo).

Please categorize your PR by adding `[Frontend]` or `[Backend]` to your PR title on GitHub.
## License
See [LICENSE](https://github.com/lerichardson/tracking-demo/blob/main/LICENSE).