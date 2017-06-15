package globalidentity

var ResponseProcessor ResponseProcessable 

func init () {
	ResponseProcessor = new(responseProcessor)
} 