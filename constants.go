package covidAss2

import "time"

// Url Paths
const DEFAULT_PATH = "/"

//const COLLECTION_PATH 		= "/collection"
const COVIDTRACKER = "https://covidtracker.bsg.ox.ac.uk/about-api"
const COVIDGRAPHQL = "https://covid19-graphql.vercel.app/"
const COVIDCASES = "corona/v1/cases/"
const COVIDPOLICY = "corona/v1/policy/"
const COVIDSTATUS = "corona/v1/status/"
const COVIDNOTIFICATIONS = "corona/v1/notifications/"
const RESOURCE_ROOT_PATH = "http://localhost:8080/"

// Global timer
var Start time.Time
