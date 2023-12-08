package main

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mainSuite struct {
	suite.Suite
}

func Test_mainSuite(t *testing.T) {
	suite.Run(t, &mainSuite{})
}

func (s *mainSuite) Test_main1() {
	require.Equal(s.T(), 1*765+2*684, Solve2([]string{"32T3K 765", "T55J5 684"}))
}

func (s *mainSuite) Test_main2() {
	require.Equal(s.T(), 1*765+2*28+3*684, Solve2([]string{"32T3K 765", "T55J5 684", "KK677 28"}))
}

func (s *mainSuite) Test_main3() {
	require.Equal(s.T(), 1*765+2*28+3*684+4*220, Solve2([]string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220"}))
}

func (s *mainSuite) Test_main4() {
	require.Equal(s.T(), 1*765+2*28+3*684+4*483+5*220, Solve2([]string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}))
}

func (s *mainSuite) Test_main5() {
	require.Equal(s.T(), 1*640+2*324+3*595+4*259+5*656+6*2, Solve2([]string{"55559 2", "Q525J 656", "457J4 324", "T6668 259", "38847 640", "777T3 595"}))
}

func (s *mainSuite) Test_main6() {
	require.Equal(s.T(), 1*28+2*648+3*723+4*7+5*282+6*736, Solve2([]string{"J44KJ 282", "5J55J 736", "T9ATT 648", "J22TJ 7", "2T7TQ 28", "38338 723"}))
}

func (s *mainSuite) Test_main7() {
	require.Equal(s.T(), 1*907+2*448+3*893+4*706+5*967, Solve2([]string{"JTTQ5 907", "JQ866 448", "6682J 706", "JT9T9 967", "3JQ8J 893"}))
}

func (s *mainSuite) Test_main8() {
	require.Equal(s.T(), 2*270, Solve2([]string{"J1234 0", "JJJJJ 270"}))
}
