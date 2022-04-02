package main

import (
	"log"
	"time"
	"twitter-reply-bot/client"
	"twitter-reply-bot/process"

	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	tweetTexts := [...]string{
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073614501405786113/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073617799940669441/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073615600917413889/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073612302382530561/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073616700429041665/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073611202870902785/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073610103359275009/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073607904336019457/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073609003847647233/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073605705312763905/",
		"#nfts #nft #nftart #nftcommunity #nftcollector #nftartist #cryptoart #digitalart #art @nft__space @NftsPlanett @NFTGalIery @Nfttender @cryptopunksnfts @opensea @BoredApeYC @doodles Check out this item on OpenSea https://opensea.io/assets/matic/0x2953399124f0cbb46d2cbacd8a89cf0599974963/28335066972685040560970427420027460738854532965002217900673073606804824391681/",
	}
	variables := process.Variables()
	creds := client.Credentials{
		AccessToken:       variables.ACCESS_TOKEN,
		AccessTokenSecret: variables.ACCESS_TOKEN_SECRET,
		ConsumerKey:       variables.CONSUMER_KEY,
		ConsumerSecret:    variables.CONSUMER_SECRET,
	}

	client, err := client.GetClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	for i := 0; i < len(tweetTexts); i++ {
		tweet, resp, err := client.Statuses.Update(tweetTexts[i], &twitter.StatusUpdateParams{})
		if err != nil {
			log.Println(err)
		}

		log.Printf("%+v\n", resp)
		log.Printf("%+v\n", tweet)
		time.Sleep(8 * time.Second)
	}

	// tweet, resp, err := client.Statuses.Update("A Test Tweet from a new Bot I'm building!", nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("%+v\n", resp)
	// log.Printf("%+v\n", tweet)
}
