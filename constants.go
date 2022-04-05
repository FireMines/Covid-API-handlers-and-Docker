package covidAss2

import "time"

// Url Paths
const DEFAULT_PATH = "/"
const VERSIONNUMBER = "v1"

//const COLLECTION_PATH 		= "/collection"
const COVIDGRAPHQL = "https://covid19-graphql.vercel.app/"
const COVIDTRACKER = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/"
const COVIDTRACKERCHECK = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/NOR/2018-09-17"
const COVIDCASES = "/corona/" + VERSIONNUMBER + "/cases/"
const COVIDPOLICY = "/corona/" + VERSIONNUMBER + "/policy/"
const COVIDSTATUS = "/corona/" + VERSIONNUMBER + "/status/"
const COVIDNOTIFICATIONS = "corona/" + VERSIONNUMBER + "/notifications/"
const RESOURCE_ROOT_PATH = "http://localhost:8080/"

// Global timer
var Start time.Time
