// Package apps provides a clean way for Tidbyt to be able to get a list of all
// community apps.
package apps

// Code generated by tools/generator. DO NOT EDIT.

import (
	"errors"

	"tidbyt.dev/community/apps/abstractclock"
	"tidbyt.dev/community/apps/afl"
	"tidbyt.dev/community/apps/ambientweather"
	"tidbyt.dev/community/apps/analogclock"
	"tidbyt.dev/community/apps/analogtime"
	"tidbyt.dev/community/apps/arcadeclassics"
	"tidbyt.dev/community/apps/baywheels"
	"tidbyt.dev/community/apps/bgghotness"
	"tidbyt.dev/community/apps/biblevotd"
	"tidbyt.dev/community/apps/bigclock"
	"tidbyt.dev/community/apps/binaryclock"
	"tidbyt.dev/community/apps/burgeroftheday"
	"tidbyt.dev/community/apps/charlestownferry"
	"tidbyt.dev/community/apps/chessviewer"
	"tidbyt.dev/community/apps/climateclock"
	"tidbyt.dev/community/apps/clockbyhenry"
	"tidbyt.dev/community/apps/cltlightrail"
	"tidbyt.dev/community/apps/coingeckoprice"
	"tidbyt.dev/community/apps/countdownclock"
	"tidbyt.dev/community/apps/cryptotracker"
	"tidbyt.dev/community/apps/dailykanji"
	"tidbyt.dev/community/apps/dateprogress"
	"tidbyt.dev/community/apps/datetimeclock"
	"tidbyt.dev/community/apps/daynightmap"
	"tidbyt.dev/community/apps/destiny2stats"
	"tidbyt.dev/community/apps/digibyteprice"
	"tidbyt.dev/community/apps/digitalrain"
	"tidbyt.dev/community/apps/duolingo"
	"tidbyt.dev/community/apps/dutchfuzzyclock"
	"tidbyt.dev/community/apps/dvdlogo"
	"tidbyt.dev/community/apps/dwheadline"
	"tidbyt.dev/community/apps/emojilingo"
	"tidbyt.dev/community/apps/espnnews"
	"tidbyt.dev/community/apps/fairfaxconnector"
	"tidbyt.dev/community/apps/fantasynamegen"
	"tidbyt.dev/community/apps/finevent"
	"tidbyt.dev/community/apps/fishbyt"
	"tidbyt.dev/community/apps/fitbitweight"
	"tidbyt.dev/community/apps/fivesomewhere"
	"tidbyt.dev/community/apps/flags"
	"tidbyt.dev/community/apps/formula1"
	"tidbyt.dev/community/apps/fullybinarytime"
	"tidbyt.dev/community/apps/fuzzyclock"
	"tidbyt.dev/community/apps/gapilotbuddy"
	"tidbyt.dev/community/apps/goodservice"
	"tidbyt.dev/community/apps/happyhour"
	"tidbyt.dev/community/apps/hurricanetracker"
	"tidbyt.dev/community/apps/hvvdepartures"
	"tidbyt.dev/community/apps/ifparank"
	"tidbyt.dev/community/apps/indegostations"
	"tidbyt.dev/community/apps/isitchristmas"
	"tidbyt.dev/community/apps/islamicprayer"
	"tidbyt.dev/community/apps/isstracker"
	"tidbyt.dev/community/apps/jokesjokeapi"
	"tidbyt.dev/community/apps/kickstarter"
	"tidbyt.dev/community/apps/kielferry"
	"tidbyt.dev/community/apps/launchcountdown"
	"tidbyt.dev/community/apps/life"
	"tidbyt.dev/community/apps/lirr"
	"tidbyt.dev/community/apps/londonbusstop"
	"tidbyt.dev/community/apps/manifest"
	"tidbyt.dev/community/apps/mbta"
	"tidbyt.dev/community/apps/mbtanewtrains"
	"tidbyt.dev/community/apps/metar"
	"tidbyt.dev/community/apps/mindthegap"
	"tidbyt.dev/community/apps/mlbleaders"
	"tidbyt.dev/community/apps/mlbscores"
	"tidbyt.dev/community/apps/mlsscores"
	"tidbyt.dev/community/apps/mnlightrail"
	"tidbyt.dev/community/apps/moretransit"
	"tidbyt.dev/community/apps/moviequotes"
	"tidbyt.dev/community/apps/natdex"
	"tidbyt.dev/community/apps/nationaltoday"
	"tidbyt.dev/community/apps/nbascores"
	"tidbyt.dev/community/apps/nearearthobjs"
	"tidbyt.dev/community/apps/netatmo"
	"tidbyt.dev/community/apps/nflscores"
	"tidbyt.dev/community/apps/nft"
	"tidbyt.dev/community/apps/nhllive"
	"tidbyt.dev/community/apps/nhlnextgame"
	"tidbyt.dev/community/apps/nhlscores"
	"tidbyt.dev/community/apps/nightscout"
	"tidbyt.dev/community/apps/nixelclock"
	"tidbyt.dev/community/apps/noaabuoy"
	"tidbyt.dev/community/apps/nyancat"
	"tidbyt.dev/community/apps/nycbus"
	"tidbyt.dev/community/apps/ogsgamesviewer"
	"tidbyt.dev/community/apps/ohhighwaysigns"
	"tidbyt.dev/community/apps/pagerduty"
	"tidbyt.dev/community/apps/pathtrainschedule"
	"tidbyt.dev/community/apps/petpikachu"
	"tidbyt.dev/community/apps/phaseofmoon"
	"tidbyt.dev/community/apps/pokedex"
	"tidbyt.dev/community/apps/pollencount"
	"tidbyt.dev/community/apps/powerball"
	"tidbyt.dev/community/apps/preciousmetals"
	"tidbyt.dev/community/apps/pubgstats"
	"tidbyt.dev/community/apps/purpleair"
	"tidbyt.dev/community/apps/randomslackmoji"
	"tidbyt.dev/community/apps/redditimages"
	"tidbyt.dev/community/apps/redditrplace"
	"tidbyt.dev/community/apps/sbbtimetable"
	"tidbyt.dev/community/apps/severewxalertsusa"
	"tidbyt.dev/community/apps/sfnextmuni"
	"tidbyt.dev/community/apps/shopifychart"
	"tidbyt.dev/community/apps/shuffleimages"
	"tidbyt.dev/community/apps/snyk"
	"tidbyt.dev/community/apps/soundtransit"
	"tidbyt.dev/community/apps/sportsrankings"
	"tidbyt.dev/community/apps/sportsscores"
	"tidbyt.dev/community/apps/sportsstandings"
	"tidbyt.dev/community/apps/spotthestation"
	"tidbyt.dev/community/apps/stateflags"
	"tidbyt.dev/community/apps/statuspage"
	"tidbyt.dev/community/apps/steam"
	"tidbyt.dev/community/apps/stepcounter"
	"tidbyt.dev/community/apps/stockticker"
	"tidbyt.dev/community/apps/strava"
	"tidbyt.dev/community/apps/subreddit"
	"tidbyt.dev/community/apps/sunrisesunset"
	"tidbyt.dev/community/apps/supermariokart"
	"tidbyt.dev/community/apps/surfforecast"
	"tidbyt.dev/community/apps/surflive"
	"tidbyt.dev/community/apps/tartan"
	"tidbyt.dev/community/apps/tcatbusarrivals"
	"tidbyt.dev/community/apps/tempest"
	"tidbyt.dev/community/apps/testpatterns"
	"tidbyt.dev/community/apps/theysaidso"
	"tidbyt.dev/community/apps/tindiesales"
	"tidbyt.dev/community/apps/todoist"
	"tidbyt.dev/community/apps/todoistnext"
	"tidbyt.dev/community/apps/traffic"
	"tidbyt.dev/community/apps/transsee"
	"tidbyt.dev/community/apps/tube"
	"tidbyt.dev/community/apps/tubestatus"
	"tidbyt.dev/community/apps/tvquotes"
	"tidbyt.dev/community/apps/twitch"
	"tidbyt.dev/community/apps/twitterfollows"
	"tidbyt.dev/community/apps/unsplash"
	"tidbyt.dev/community/apps/usgsearthquakes"
	"tidbyt.dev/community/apps/usyieldcurve"
	"tidbyt.dev/community/apps/vergetaglines"
	"tidbyt.dev/community/apps/verticalmessage"
	"tidbyt.dev/community/apps/wantedposter"
	"tidbyt.dev/community/apps/warframecycles"
	"tidbyt.dev/community/apps/weathermap"
	"tidbyt.dev/community/apps/web3counter"
	"tidbyt.dev/community/apps/whosthatpokemon"
	"tidbyt.dev/community/apps/wordlebyt"
	"tidbyt.dev/community/apps/wordoftheday"
	"tidbyt.dev/community/apps/worldclock"
)

// GetManifests returns a list of all apps in the this repository. Add your applet
// below to include it in the Tidbyt Mobile app for all Tidbyt users.
func GetManifests() []manifest.Manifest {
	return []manifest.Manifest{
		abstractclock.New(),
		afl.New(),
		ambientweather.New(),
		analogclock.New(),
		analogtime.New(),
		arcadeclassics.New(),
		baywheels.New(),
		bgghotness.New(),
		biblevotd.New(),
		bigclock.New(),
		binaryclock.New(),
		burgeroftheday.New(),
		charlestownferry.New(),
		chessviewer.New(),
		climateclock.New(),
		clockbyhenry.New(),
		cltlightrail.New(),
		coingeckoprice.New(),
		countdownclock.New(),
		cryptotracker.New(),
		dailykanji.New(),
		dateprogress.New(),
		datetimeclock.New(),
		daynightmap.New(),
		destiny2stats.New(),
		digibyteprice.New(),
		digitalrain.New(),
		duolingo.New(),
		dutchfuzzyclock.New(),
		dvdlogo.New(),
		dwheadline.New(),
		emojilingo.New(),
		espnnews.New(),
		fairfaxconnector.New(),
		fantasynamegen.New(),
		finevent.New(),
		fishbyt.New(),
		fitbitweight.New(),
		fivesomewhere.New(),
		flags.New(),
		formula1.New(),
		fullybinarytime.New(),
		fuzzyclock.New(),
		gapilotbuddy.New(),
		goodservice.New(),
		happyhour.New(),
		hurricanetracker.New(),
		hvvdepartures.New(),
		ifparank.New(),
		indegostations.New(),
		isitchristmas.New(),
		islamicprayer.New(),
		isstracker.New(),
		jokesjokeapi.New(),
		kickstarter.New(),
		kielferry.New(),
		launchcountdown.New(),
		life.New(),
		lirr.New(),
		londonbusstop.New(),
		mbta.New(),
		mbtanewtrains.New(),
		metar.New(),
		mindthegap.New(),
		mlbleaders.New(),
		mlbscores.New(),
		mlsscores.New(),
		mnlightrail.New(),
		moretransit.New(),
		moviequotes.New(),
		natdex.New(),
		nationaltoday.New(),
		nbascores.New(),
		nearearthobjs.New(),
		netatmo.New(),
		nflscores.New(),
		nft.New(),
		nhllive.New(),
		nhlnextgame.New(),
		nhlscores.New(),
		nightscout.New(),
		nixelclock.New(),
		noaabuoy.New(),
		nyancat.New(),
		nycbus.New(),
		ogsgamesviewer.New(),
		ohhighwaysigns.New(),
		pagerduty.New(),
		pathtrainschedule.New(),
		petpikachu.New(),
		phaseofmoon.New(),
		pokedex.New(),
		pollencount.New(),
		powerball.New(),
		preciousmetals.New(),
		pubgstats.New(),
		purpleair.New(),
		randomslackmoji.New(),
		redditimages.New(),
		redditrplace.New(),
		sbbtimetable.New(),
		severewxalertsusa.New(),
		sfnextmuni.New(),
		shopifychart.New(),
		shuffleimages.New(),
		snyk.New(),
		soundtransit.New(),
		sportsrankings.New(),
		sportsscores.New(),
		sportsstandings.New(),
		spotthestation.New(),
		stateflags.New(),
		statuspage.New(),
		steam.New(),
		stepcounter.New(),
		stockticker.New(),
		strava.New(),
		subreddit.New(),
		sunrisesunset.New(),
		supermariokart.New(),
		surfforecast.New(),
		surflive.New(),
		tartan.New(),
		tcatbusarrivals.New(),
		tempest.New(),
		testpatterns.New(),
		theysaidso.New(),
		tindiesales.New(),
		todoist.New(),
		todoistnext.New(),
		traffic.New(),
		transsee.New(),
		tube.New(),
		tubestatus.New(),
		tvquotes.New(),
		twitch.New(),
		twitterfollows.New(),
		unsplash.New(),
		usgsearthquakes.New(),
		usyieldcurve.New(),
		vergetaglines.New(),
		verticalmessage.New(),
		wantedposter.New(),
		warframecycles.New(),
		weathermap.New(),
		web3counter.New(),
		whosthatpokemon.New(),
		wordlebyt.New(),
		wordoftheday.New(),
		worldclock.New(),
	}
}

// FindManifest finds an manifest at the given ID.
func FindManifest(id string) (*manifest.Manifest, error) {
	for _, app := range GetManifests() {
		if app.ID == id {
			return &app, nil
		}
	}

	return nil, errors.New("app manifest does not exist")
}