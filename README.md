# IFTTT Glue

An example API-middleware/SMS approach to chaining IFTTT recipes conditionally.

### Setup

* Get the [Google App Engine SDK (for golang).](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go)
* Clone this repo into your `/path/to/go/src` directory
* Copy `config.go.example` to `config.go`
* Enter your Twilio/IFTTT info into `config.go`
* Run it! `$ goapp serve`

### Background

The Google OnHub router just got a new IFTTT integration. Google/IFTTT challenged my brother and I to come up with a as many integrations as we could using the OnHub and a huge box of smart-home stuff.

See the video here: [https://www.youtube.com/watch?v=JPH74ZHDuCI](https://www.youtube.com/watch?v=JPH74ZHDuCI)

With IFTTT you can trigger recipes from the OnHub when a) a device joins the network, or b) when a device leaves the network. We cooked up a lot of basic recipes around that premise. But we were also interested in using the OnHub as a kind of situational monitor to do more complex recipes like:

> IF the Nest camera senses motion, AND nobody is on the router network, THEN sound the D-Link siren!

IFTTT doesn’t support this kind of contextual chaining out of the box. But we found you can patch independent recipes together with a little middleware magic. So, the above recipe becomes:

> IF kylePhone connects to the OnHub, THEN post request to https://ifttt-glue.appspot.com/onhub/connect/kylePhone

> IF kylePhone disconnects to the OnHub, THEN post request to https://ifttt-glue.appspot.com/onhub/disconnect/kylePhone

> IF brendanPhone connects to the OnHub, THEN post request to https://ifttt-glue.appspot.com/onhub/connect/brendanPhone

> IF brendanPhone disconnects to the OnHub, THEN post request to https://ifttt-glue.appspot.com/onhub/disconnect/brendanPhone

> IF Nest detects motion, THEN post request to https://ifttt-glue.appspot.com/nest/motion-detected

> IF SMS received with #alarm in body, THEN sound the D-Link siren!

This API will keep track of who's on or off the network, and conditionally send an SMS trigger when nobody is home... to sound the alarm!

I discuss the approach a little more in a blog post here: [https://medium.com/sea-salt-ventures/google-onhub-ifttt-81ebc2e60c24](https://medium.com/sea-salt-ventures/google-onhub-ifttt-81ebc2e60c24#.deiqipo6b).
