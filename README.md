# HumanityCoins

<!--- [![Deploy to Bluemix](https://bluemix.net/deploy/button.png)](https://bluemix.net/deploy?repository=https://github.com/vpaprots/HumanityCoins.git) --->


# Application background

This application will demonstrate a HumanityCoins point system where users can honour each other with HumanityCoins. Later HumanityCoins can also be collected by wearable devices, or smart meters, rewarding people saving resources (electricity, water, gas) or living a healthy lifestyle reported by their wearable device. To avoid abuse of the system these points are audited by a tweeting auditor (using twitter). The HumanityCoins points can be used by companies and government bodies to reward people doing good to their communities, health and to the environment.

### NOTE: This version is compatible with Hyperledger v1.0 docker containers
# Some technical details:

Types of thanks:

	1.small =1 HumanityCoins 
	2.medium=5 HumanityCoins
	3.large =10 HumanityCoins
	
Attributes of a user:

	1. userID (unique string, used as primary key)
	2. balance (int, sum of the calculated points from the type of thanks received)
	3. thanklist (string array of the thanks recieved)

Attributes of a thank:

	1.Thanker (the name of the person giving the thank)
	2.ThankType (type of the thank small/ medium/ large) 
	3.message (a small message explaining the thank, can be empty)

# Project setup:

cd $GOPATH/go/src/github.com/szlaci83/HumanityCoins/ && dep ensure

cd $GOPATH/src/github.com/szlaci83/HumanityCoins/fixtures/ && docker-compose up -d && cd .. && ./HumanityCoins 
