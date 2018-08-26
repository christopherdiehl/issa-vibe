package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var tmpl = make(map[string]*template.Template)
var genreTrackMap = map[string][]string{
	"Punk":    []string{"2TfSHkHiFO4gRztVIkggkE","1FMm3wG5WOhi7js35KR7Ea","7lRlq939cDG4SzWOF4VAnd","6SpLc7EXZIPpy0sVko0aoU","4bPQs0PHn4xbipzdPfn6du","4wzjNqjKAKDU82e8uMhzmr","3ZsexY07D4t4HRq7ogeSAS","6L89mwZXSOwYl76YXfX13s","0a7BloCiNzLDD9qSQHh5m7","2D7tauy2bntJnJQ2C4rO4x","7dN52YlyGl6c9tmxdv6zMp","4rLKjYqJjLU0Odg52kGL2O","5lDriBxJd22IhOH9zTcFrV","0BRHnOFm6sjxN1i9LJrUDu","47No93LxERvV6MtOAmQzHS","1fJFuvU2ldmeAm5nFIHcPP","4c7xQ7OsqL6W4RwEQmID7g","7wOmQJeVX6qjNXqqsKOkPx","23oxJmDc1V9uLUSmN2LIvx","33iv3wnGMrrDugd7GBso1z","0wVluBsVAVzBKrqspuCcwR","3u7EIgAlwNQtxkM4bOA7uI","78OBZL4Z6QHdwwloWIh0Lt","2g2a5kDeZexbUTD8abcvm6","3asFGFY3uLjMDmML1p0tYm","12mF5rnbbT7jNqiBv8NBFt","45zvStEMsXp8z45OQRhWFJ","27L8sESb3KR79asDUBu8nW","4jVhSbMMHctghoOTFuLPB2","1iZTuY3wspO4DxFXBD58WX","1LkoYGxmYpO6QSEvY5C0Zl","3VA4sjTMSTTF02hFGmlpJh","0sNKiz82ATCvT3f3XVVUUj","4UuHWIkAWnN176A2rx7YFB","3tWudnj1NK7YUfnR91nezW","5GhYfK0jevTEtXOaqCKCxw","1hIupFeRu3nmMcjbjxPnMc","2jxujpjjRkggQbPPNhVseJ","1Dr1fXbc2IxaK1Mu8P8Khz","0Ti2dlF2xLjXblvdU5fCxM","3C0f4W0lQSgMARICJF3KLh","0ChpoNjXfJPjMvCIN6so6J","6vFD3c1WI1zuKoyk22dMw3","5BmagRD7Thki6O1zZwbxBy","38z6jaLE0F13dxv28n3wlq","0DKNNR9iDjwfCEpMiFXMJq","4sQJXbtxwoKHG7PwMRJ0Ig","1aYWQrIT8XPdtV4hhUjczN","6oiaq7aC6DooFHBFJkmQ1n","4cqYUqmKMcb3q1tdImVpGW"},
	"Jazz":    []string{"6J8Z1UJG7uTS1mhZsgmJCp","1z2FcRqCR6byfXDSvHq8E4","5FXIXX9FniNJfHd9DJrOop","78RAfPGFpTUqnPSTV0LcmT","2ncnr8AuzqAabOGEIoA863","2seZ8ZHnD5RQBkhT8RsBZu","6vZ8PZZy7ljlQEgCPSZTAM","4Hus6AQNHezUdbBKuo1aE8","4gCaFMkA1CERmgy1ny9fhy","50ma4LjaP5XDL1BsHrZ1D9","0D5PNTs1HSYe2nHlHulJQw","3ZnPVp4Frub6UadiiiGDD9","3LZ0GUgBO1XXW6jKK5A2lJ","79MOPpvNSagTPcwnx4Ct2b","49e8jyxGYok1DJRLqOlg4T","76Bk0cDFzg1ljHQpY0SUqV","2SzfUlXBgyYKTJMPbbqPw7","3MU1qPFFyJQy1RbmpLH0mb","1XrEY8nz2P4L9rVoq9MdXs","2ovibjyd0xaW7HvfGvx0Qe","7kdZlgkWHnivQNSGe6AsTa","6BMmyctJWWdbTHKpXkrja0","59LTHdhW1H688shHvcdub5","42q9VqB8ooQE5yn3mup3h8","5fhLDCuNimt2zvkiFxOAbU","1Rx7TVIX8t9pZFJdLq0vvY","5adNkxFBgFYu3N6m8J6muG","0Y6WjYBWld3reWsTMT8FTd","49vQNdC75cHdQ6JLsxGKPJ","4JHItqtk4k5wv207UO1D66","75MUAXmFXalRbe44eYPzfG","3s2RFp5hU6jEvAmfZrnrAi","4Uecyf82f85MdvnGLA4SWE","2r3dQilQpgjPbRomN7GAvD","4z6d79RAbC5q0VAqPQas8W","2yuDexdINZo3jxbFHf5fLw","1iNE14jpYWtLoeQitCoYaG","5zpqkOpg31ksUdhuqJ9CQU","1zowWjoCroDC3WZYxbrlNQ","5Bl1TC8Y1qh3ui6fGSWL5a","58TwX5CF328t6qDm9sgaLb","12zuJhwGaqJu8fvgBWqYuy","4ZUlQw3KS7qiBrgJxSyel0","3odLwNmrGr91br8fhtrP1v","6bqNdzv1KHd5C7ae0YbV7r","3lw4OExcjj4kba4uU9gF4P","7u38kj1yGDoExUnP6gxxor","1Kez08abn4yr3T7hbCeW3i","3PJMsxg6rz9FOo6xNiASXz","5bjV2B1rH2gkdtJ4qr0qQy","0Uu93Y0BiC2h01Al6QpvWY","79fmhV6NAraE4Xz3MIXnN4","3ONcanYtZHMVcgqPDAsVHd","0Ya7lMXJGmETT1G3bW0uX2","2Smeb5R5N2aCYAi7cIhema","4N7iWwgqcWr3p5JzngOSEF","70S7sASwrZH3Fass66MKqJ","3dtrI8OOHWxMyg6EHNecdt","7ysmJhXFQtiBQlk6EZ6sks","499pC27JxKGMh6mYyuRmlT","6f6OQJv9qjqyJq9NHin45n","3MhWFpRMVEyZ9QXGIb5RJK","0aWMVrwxPNYkKmFthzmpRi","6NpwgYdHNNpq81GnpJNx3Y","2JR1WeETsXmFR1kRC2oXXq","14aFuIi1l2arBrZfBv9fBP","58VcgWixvxnMdt4bj29PdQ","75A0vJQiBzAxkC7pMjoSeX","4zRweMHazlqCTUe6GvviYf","4vuoL9ldmpmZcXm7RkhQvx","1axzRYUUASTfN4OOtrjfQd","1AYtj1r4CcTiU6qMcmy0Hz","1J9iVCaEriyoMXpj2XszhU","4lj4cW45UP5tQ8AhRxnQi9","4zPDjRgceGrR0Sm54YDTTf","1PMlAs24OBYlwXF7OeiVYg","1ohoaGnztWcfczuRDceRdn","2AtLQju6O4Slbhtuh2ndCD","4G3EKvGtqcUSZqPQIWJFxd","2sZ1zWeLaquu9tx8IUMuh7","1UceH3dNaxzKSG4qQ1kbp2","5wRlH3diPDfn9jv91hxkUy","5LaTsDrhrpXwLniidLpnF7","3wPZVnOScmpbn6JkWJxi3S","0cZuS1VcwpKi5BxP4ESDaj","082zyi189u5lRp2gPkE8Lb","5qRJZGqWHMpKRxk8rmAVto","4bBIADw4urJPkcx0YOfEIo","4ymHy4hzJ09WxvvT7p0Azy","24M4rafgMYq9b49CEYPpm9","7hbpUZY7qsDPBe9X2juJEz","0gPxAMKPYhnFSAKkBcy3UN","4UhBGmqcApvAZFTq53ukh7","2oFgZSzpbXscqwZQyoFIzW","5bF6crgJa5Q58IWgbdUbuO","2R9KveaVtFgiiEC6KSF9un","6rK27cC9hX0e1Chrnxo9Vn","0WPHL9IRv1Nm4xu21qxgnw","0tiIQkcdPcECxqbS9dgDXx","6kXSVoByQngZxSqX4kSGgF"},
	"Blues":   []string{"7g4fX37Y3lziLMoxrTTGI3","5yeM63cXgGvSN2VrcHbv6x","2DBFAJgsqhYk5Z1AF7tAMH","4pPWNeApSiORRQucFZt85y","5yuqWMCOtMY0IBaQCBzqT5","1BQkVDlEOtYKOaotnJuNzz","3FYVkmrVYbEo17vXmb5x2A","0lwUIYJfWGB8qUEAJtYfb6","4Vc3bzFxXdrABa2DxgyCMT","5A6t8zfY6Uj6fxOYOYppJK","0FbKvqTBo2TsSzhT5ohFI2","3UnOgOZkTRvlmfUcTHzvc6","1EXyAPbplzZKhFZx54smQW","2b2oHgRMo5QYRLft2TRDFi","1Yc25rbbGJZCtz8dtizMKY","2BEqkiPEzz4tYul7V0SGxW","00AVBnaQz7yRv8NTpNlCVU","0VVBiLcDZthpwsXXBCXPgu","47dOtJjjwVkbKW244db2dE","06PsyBBc0YCQhGNlfPZ1SS","7b1PDaZvnxGParUPi9ViRJ","1zHHwTIloBthaYONQX2TZ0","5WKhPe4lSUYgHhhwFMnX4Y","1YrdxppVdO7a5S99gudh43","27B8NcE66RUKHP3cN46Fhx","37C9PC2PjE2oURNjOyzof4","0lcsIpA9Ff2DZufaMIQe4k","1OtGo99uypkRbMqshBVFnn","32G6qOwK5SuBEA9D6pth34","03HbxaCm8eBHLxbfdRRQ2e","5VB9VqBUzxR8lj49Scmly0","3DDmBaZNrxKbGApdrYXBAc","5QJMMpa8ihBI2CFEhhmOi4","1gpyfbutOl79qgVjMrZQOe","257L4vqFRw2IG67CkHQ0ff","2i9hkOptEKySb9AmLkWg1w","6AvLsGpRX9GEej0zB7evfj","54jlEZOJZA8TwEGnKLYFon","3T76zPJz3tWL27FrjJe2ot","6RjnBLAqjop8cgSsPrUlCs","35upu7uVMA4XNiQMMl6ZGh","4DnK27URnJcCULRNDHvm2M","3ZTScjTJav3Yak0OqLTuAL","3eu6yCblDDyuqvDTkPZksu","4pF0B8Fae3dQrsuWds2iQ9","57XbiF3mXSAXDntyEw57lW","1LfGP4l3u94cu57cR4khzq","3WlKwc0STwk7MYHyn2FjXm","39EQX63g5VOvBOpcN1orlQ","5XYmnUaaZ3D9Sqdt9SInEn","7FqrsV0vBwNiQNQI6jfzni","0jVfp8deBjwdoi2oEm5hgK","2pDOP7lUpO5LGUz798H0me","0F3ZS2gMHCImgws6ewz1H7","5ke9Y5BL0wbGOg4pshd1B0","5aGkIyzLVYGuLIHkxpNAFF","4z9Je08Qpyfn6q7bbiqmW5"},
	"Pop":     []string{"4OafepJy2teCjYJbvFE60J","4keoy2fqgwGnbWlm3ZVZFa","0d2iYfpKoM0QCKvcLCkBao","5WvAo7DNuPRmk4APhdPzi8","4VUwkH455At9kENOfzTqmF","21RzyxY3EFaxVy6K4RqaU9","3EPXxR3ImUwfayaurPi3cm","0E9ZjEAyAwOXZ7wJC0PD33","2dpaYNEQHiRxtZbfNsse99","5w0IaaNTTnrXvuwJR8L87d","2xLMifQCjDGFmkHkpNLD9h","3u1S1OmAUhx5DRlLrXqyp3","5N5k9nd479b1xpDZ4usjrg","6FRLCMO5TUHTexlWo8ym1W","0WdR2AyLW1Drd3OUdwezM0","0bAkKNCQfWkexHFn7fIKns","2lYTJK94hb0fd1LQtb6Dhk","55S2PQgSMYAhgoTCcGCDfw","04ZTP5KsCypmtCmQg5tH9R","5IaHrVsrferBYDm0bDyABy","4ofwffwvvnbSkrMSCKQDaC","7gkWXbAxIYuvtOpcN3p9GJ","4kWO6O1BUXcZmaxitpVUwp","6AkFxZzBfyzV5RF7mGkASM","4QVS8YCpK71R4FsxSMCjhP","5a8RPWgKSmcGBGcffmIrUi","5yNEdBlZMxpcVtGKz5NFk5","4zIO3ilp5HvTeK3HJHxhMP","4ut5G4rgB1ClpMTMfjoIuy","6kPJZM97LwdG9QIsT7khp6","0OgGn1ofaj55l2PcihQQGV","6X6SdJzxyw7GzxwP2fr9WV","1rqqCSm0Qe4I9rUvWncaom","1snWlbcbgQpJfknoI30DWG","55GiOwtxNajXVbpjjIzMnD","59D6pNWyn8UXPpos5fbAqm","5ubwXXOsH9bUhZsQ0Cybha","3T4UodGkfZObJ43RtA5KFU","58q2HKrzhC3ozto2nDdN4z","2E124GmJRnBJuXbTb4cPUB","0p0ljM6RxgpGt7wthGqBZa","6XcfKZvJio9Z0fQy11GnNX","32iYwQ4OYiurnGnOfJ1aEh","4txpZk7WSjV1dsZAw5WYcT","4gB61mP4tNChn4LgfquMhi","6vsV4D8BM6PioRr1UOx0n2","63SevszngYpZOwf63o61K4","56gH0zGe3SP9SGNoQ2pWwn","4Pbg79cTBu4vgSphoyNq3j","7DrluKkTviBwCc8AV3VGmf"},
	"Rock":    []string{"0hCB0YR03f6AmQaHbwWDe8","08mG3Y1vljYA6bvDt4Wqkj","2aoo2jlRnM3A0NyLQqMN2f","39shmbIHICJ2Wxnk1fPSdz","1RJeiAIwR9pZBgJA8ndZLL","66lOpKgTyFjOrac4S1s94g","3YBZIN3rekqsKxbJc9FZko","7MRyJPksH3G2cXHN8UKYzP","57JVGBtBLCfHw2muk5416J","0iA3xXSkSCiJywKyo1UKjQ","5MxNLUsfh7uzROypsoO5qe","61Q9oJNd9hJQFhSDh6Qlap","4CJVkjo5WpmUAKp3R44LNb","4cCQ0wBIc9YHhvOWvWNj7S","2aSFLiDPreOVP6KHiWk4lF","2LawezPeJhN4AWuSB0GtAU","6cr6UDpkjEaMQ80OjWqEBQ","70YvYr2hGlS01bKRIho1HM","2zYzyRzz6pRmhPzyfMEC8s","1AhDOtG9vPSOmsWgNW0BEY","2X6gdRlGOQgfaXU9ALUQFQ","4XizBlyqR7ZGVTX0Fyonm2","78lgmZwycJ3nzsdgmPPGNx","7o2CTH4ctstm8TNelqjb51","2sXp9Qmvc7mRaDBjBgcGGi","5e6x5YRnMJIKvYpZxLqdpH","1IqFh00G2kvvMm8pRMpehA","54eZmuggBFJbV7k248bTTt","3k9i7UzeSUYWIfUZFeFDUd","2d4e45fmUnguxh6yqC7gNT","7GonnnalI2s19OCQO1J7Tf","1JkZg3eMQTmTn93E8Yd3UL","2HXixVqzzm9rEUIMAWzshl","0dOg1ySSI7NkpAe89Zo0b9","1xShPgQbOUa98avWJQFDBY","2EqlS6tkEnglzr7tkKAAYD","4DMKwE2E2iYDKY01C335Uw","4gMgiXfqyzZLMhsksGmbQV","5LNiqEqpDc8TuqPy79kDBu","1Ly0ssAgeM7YqdHptao8Oe","6NxsCnLeLd8Ai1TrgGxzIx","5SAUIWdZ04OxYfJFDchC7S","5tVA6TkbaAH9QMITTQRrNv","4bHsxqR3GMrXTxEPLuK5ue","7N3PAbqfTjSEU1edb2tY8j","2PzU4IB8Dr6mxV3lHuaG34","4KfSdst7rW39C0sfhArdrz","6zeE5tKyr8Nu882DQhhSQI","5p3JunprHCxClJjOmcLV8G","3YRCqOhFifThpSRFJ1VWFM","2vX5WL7s6UdeQyweZEx7PP","7LRMbd3LEoV5wZJvXT1Lwb","0J6mQxEZnlRt9ymzFntA6z","5UwbnHhjnbinJH8TefuQfN","5QTxFnGygVM4jFQiBovmRo","6gXrEUzibufX9xYPk3HD5p","32kgOw8wejH7zUhtXCM8DH","1dv3ePjze9tPq2pk8eWJdR","63OFKbMaZSDZ4wtesuuq6f","7IhkWF9zpiYKlU7CxHxYT9","4JNi40t7xR5bO3PWxRkiPN","2SiXAy7TuUkycRVbbWDEpo","3aoDEt6zSuYQ47gzarlaVo","2b9lp5A6CqSzwOrBfAFhof","0eFvoRSTTaR2q8bSWVjwfp","0F0MA0ns8oXwGw66B2BSXm","1qRA5BS78u3gME0loMl9AA","2KmEgiY8fQs0G6WNxtzQKr","5y4761pjcksXs1hNsCmc4n","54flyrjcdnQdco7300avMJ","6NTqBHONQqmud0ONBzsLfZ","5wj4E6IsrVtn8IBJQOd0Cl","1zGk1kgAxCpg6PHjGuLe4J","51H2y6YrNNXcy3dfc3qSbA","0832Tptls5YicHPGgw7ssP","2f0P7iELCvAlV8j6Z3rGDE","38Ngied9rBORlAbLYNCl4k","0QwZfbw26QeUoIy82Z2jYp","3sY6z4pGcxpMwx3W026WtU","6Vjk8MNXpQpi0F4BefdTyq","6KTv0Z8BmVqM7DPxbGzpVC","3IOQZRcEkplCXg6LofKqE9","7Ar4G7Ci11gpt6sfH9Cgz5","52dm9op3rbfAkc1LGXgipW","5eYwDBLucWfWI5KsV7oYX2","3G69vJMWsX6ZohTykad2AU","5BIMPccDwShpXq784RJlJp","1lhpxZT57yw5toGJtt8fGE","5EcvzboYmDIfXo9gPplsSB","14XWXWv5FoCbFzLksawpEe","4YMLW13PuuO7o3jkIWlKAs","24NwBd5vZ2CK8VOQVnqdxr","0NWPxcsf5vdjdiFUI8NgkP","0N0q5D8KmH13mmODXJsJqS","0FeCO85RKW8fDRytwXof2x","0vSiz90JRxvYsbCTx84oVM","4Zau4QvgyxWiWQ5KQrwL43","7iLGljK0LNpGamQQEOuCRC","1oG2vWELiGjIqxwqGcyqwF","7cy1bEJV6FCtDaYpsk8aG6"},
	"Hip Hop": []string{"5ByAIlEEnxYdvpnezg7HTX","503OTo2dSqe7qk76rgsbep","3ssX20QT5c3nA9wk78V1LQ","4IYKjN1DrYzxKXt0umJqsG","0AvV49z4EPz5ocYN7eKGAK","7lQWRAjyhTpCWFC0jmclT4","7KwZNVEaqikRSBSpyhXK2j","0Zh5U48tZNeAzzLTV1CVBE","33ZXjLCpiINn8eQIDYEPTD","2lUirvUhqfBqJzUvk4tLoK","3djNBlI7xOggg7pnsOLaNm","20XdEFyaUR9C7aDIdq2OAd","17larILaMHwsprsMFnYI77","68pWLkspLFIfIPPtzyTkQy","20LxTsa6936zOlzTWqoPVt","0WKYRFtH6KKbaNWjsxqm70","2xTft6GEZeTyWNpdX94rkf","3NGT0Td7H4Is1qrlDQJxma","4ENy0Hz7IsSf8lDz3to3FH","6OtlkByM64FzhICejC8TaY","1rylJ0VXvhOzLGNm401iaI","5H0dVldtI6HJZnEXzigKfY","4INDiWSKvqSKDEu7mh8HFz","119c93MHjrDLJTApCVGpvx","2oTDOIAdsxPTE7yAp4YOcv","2bZfoRmADhlF5gMCSZ4pDS","4akCwzFRS8Syu5uFy7L1sG","1yy2DlSDtEt90d54rPDPXz","5n276uEKrEFohrt42pP8Tf","3GKL13lkM5nRc4zC1lIOrR","4G19iP1Ltd4GLknCiS2k7y","1oTHteQbmJw15rPxPVXUTv","0l08dcPEqNEUhymVBext8h","5XhkV07Vou38wnrzwURUOC","5ltdcgEflWut0osW0ZzLhM","5PQmSHzWnlgG4EBuIqjac2","0ipLnUeK5PODrTKyT3wVGT","7N1Vjtzr1lmmCW9iasQ8YO","1n86eBhlqaxWIwi5YZTUok","4lqU4WByDFKDGrHknQV6ZC","3jFTsobEff1Cz8DlIAisYY","0shK5iZQppbHPQYiy60xs9","550JPXXPgjr1xqxtKIaX6b","5KSJ9k1FYjFLnIRlJT2wF8","2OsGdZYDdTK4pxDMsK0Blv","3W6UZZEiAIywZYpNpfBplK","1NHwvBmrUje4L1dxfWnXCH","6MszGZxe52sc4owTmjp0O9","3Ti0GdlrotgwsAVBBugv0I","47wZfF4OdME3xkIPhhpSSF","2vfvGlqCB7oertO5VLE0sz","3ABG1UQTk0eLYUAeWkmfMi","65iyI1iybyv5ecsfBHSdUf","158DIbrVt4YbqNnWyRCS3P","6hWxL0tpZm0QOLNSpT6Qra","5fpizYGbi5IQoEraj6FP0R","5PbFD3AIyunz5HxocduA9e","6Kp8tQd95YPSrSUTkll2Of","30oTS7bm0aH3p7lqjEIu8q","6hxn98poTu1O4YZfafvC18","3C09tEhYKCysPOVqmJxOTN","6enXeBnLnyjfweN6VWcIoh","3iajlDW8TWwAE3gWvPRTnI","0KkrT0y1iht0tqgh9vrGd3","3Y6XWs8xMlCngyIxNOFnsp","0trHOzAhNpGCsGBEu7dOJo","1RjU7UhRICmdFPKxCucLgq","5nG6GkBGfZmhSgthtLfUuY","27AHAtAirQapVldIm4c9ZX","24vNw0Z0srb4zYXwrakw8E","7aRGb3vZGMLNpK2PEdUjdA","52QkhYa8lRm7DD6Mb8996S","6PhOzQShyz2vjELsXax7TQ","6nwAubYnLGl7F7cDuOoVLp","34MRi5GRWHFUa731JBxFEl","5f1RQ9Xq2iS2JUx9eowJov","6WceXj9vvlxVyRCkDXt4nI","4YCBBxP0AeWPpEwidahrha","0PV1TFUMTBrDETzW6KQulB","2vpvIEM9321ME9Z63nd3Rm","2xoPn7HkB4K4liMVojQQ3T","0eEXcw3JLVXcRxYrVYMy68","7nk7oS0G6qG3ofHUAKbg9F","7x9hPUtk2PNUcSwPbMWQP5","0xyR0utsJZGtHdcuBqTEVZ","7Kutf53V4k0CvU4nKu5oK6","38fJv1GvxHkHlgbERUuBCJ","2Mb3zpobD0CvJGWv6NpsPy","2sXMYMpvGVJJBIwrzNf0Lr","6ETW8Q1Xd68dp4Ck1ieE7e","6dRj771QZr0KOQVwaB6ykM","2iYM4CGkoTfsLw9MI5K1AC","2WRzpLD8qDRrxMXc63E5WJ","3Dh10QBwBWVQAMJNNhLRq4","3YXlTtlYDXI2SyofNm1ccg","2qx0FYGob348g1BdwEyMJA","22Yq5tkMN8JE7h1Bq3pWHS","0hKr166QnNZ0a37G4UO0VY","64gaTolmXymtP6AA5urDWd","2gfBV96ou2PCp0VhvddOVQ"},
	"Folk":    []string{"2D7tauy2bntJnJQ2C4rO4x"},
	"Indie":   []string{"2D7tauy2bntJnJQ2C4rO4x"},
}
https://open.spotify.com/user/spotify/playlist/?si=XKwvwyX4SX6RC8Rw4uUkHQ
func parseTemplates() {
	tmpl["getIndex"] = template.Must(template.ParseFiles("layout.html", "genreform.html"))
	tmpl["postIndex"] = template.Must(template.ParseFiles("layout.html", "songPlayer.html"))
}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tmpl["getIndex"].ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		genre := r.FormValue("genre")
		var url strings.Builder
		genreTracks := genreTrackMap[genre]
		songIndex := rand.Intn(len(genreTracks))
		// for track use https://open.spotify.com/embed/track/
		url.WriteString("https://open.spotify.com/embed/track/")
		url.WriteString(genreTracks[songIndex])
		fmt.Println(url.String())
		err := tmpl["postIndex"].ExecuteTemplate(w, "base", &struct{ URL string }{url.String()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	parseTemplates()
	http.HandleFunc("/", handleRoot)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Please navigate to localhost:8080 to view application in browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
