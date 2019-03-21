package joinelection

const (
	NotFoundElection int = iota
	ErrorGetElection
	NotFoundVoter
	ErrorGetVoter
	NotValidNationality
	NotValidAge
	ErrorGetAllCandidates
	NotFoundCandidates
	Success
)
