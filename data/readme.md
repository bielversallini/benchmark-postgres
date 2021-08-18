1. Started with plain SNO.
2. Installed ACM.

Added to sender.go line 195
```
	ioutil.WriteFile(fmt.Sprintf("./data/sno-%d.json", count), payloadBytes, 0644)
	count++
	return nil
```

I0716 19:20:43.405854   10206 sender.go:189] Sending Resources { request: 88949, add: 4542, update: 0, delete: 0 edge add: 1880 edge delete: 0 }
I0716 19:20:48.463120   10206 sender.go:189] Sending Resources { request: 72222, add: 2, update: 0, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:21:53.707983   10206 sender.go:189] Sending Resources { request: 532210, add: 24, update: 0, delete: 0 edge add: 14 edge delete: 0 }
I0716 19:21:58.726300   10206 sender.go:189] Sending Resources { request: 656485, add: 7, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:22:03.745369   10206 sender.go:189] Sending Resources { request: 927741, add: 0, update: 3, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:22:08.767004   10206 sender.go:189] Sending Resources { request: 44421, add: 6, update: 1, delete: 0 edge add: 4 edge delete: 0 }
I0716 19:22:13.786021   10206 sender.go:189] Sending Resources { request: 882329, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:22:23.820617   10206 sender.go:189] Sending Resources { request: 232201, add: 44, update: 2, delete: 0 edge add: 6 edge delete: 0 }
I0716 19:22:28.834523   10206 sender.go:189] Sending Resources { request: 212760, add: 44, update: 33, delete: 0 edge add: 6 edge delete: 0 }
I0716 19:22:33.851627   10206 sender.go:189] Sending Resources { request: 481059, add: 29, update: 1, delete: 0 edge add: 77 edge delete: 0 }
I0716 19:23:19.028363   10206 sender.go:189] Sending Resources { request: 451295, add: 13, update: 12, delete: 0 edge add: 5 edge delete: 0 }
I0716 19:23:24.047760   10206 sender.go:189] Sending Resources { request: 949260, add: 49, update: 2, delete: 0 edge add: 5 edge delete: 0 }
I0716 19:23:29.064326   10206 sender.go:189] Sending Resources { request: 502885, add: 3, update: 1, delete: 0 edge add: 5 edge delete: 0 }
I0716 19:23:34.081821   10206 sender.go:189] Sending Resources { request: 740949, add: 49, update: 3, delete: 0 edge add: 16 edge delete: 0 }
I0716 19:23:39.100427   10206 sender.go:189] Sending Resources { request: 826028, add: 85, update: 2, delete: 0 edge add: 73 edge delete: 0 }
I0716 19:23:44.118842   10206 sender.go:189] Sending Resources { request: 783267, add: 26, update: 0, delete: 0 edge add: 42 edge delete: 0 }
I0716 19:23:49.138652   10206 sender.go:189] Sending Resources { request: 685049, add: 2, update: 4, delete: 0 edge add: 1 edge delete: 0 }
I0716 19:23:54.155327   10206 sender.go:189] Sending Resources { request: 575205, add: 0, update: 2, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:23:59.175977   10206 sender.go:189] Sending Resources { request: 454076, add: 86, update: 4, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:24:04.193091   10206 sender.go:189] Sending Resources { request: 494352, add: 50, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:24:09.211300   10206 sender.go:189] Sending Resources { request: 728689, add: 0, update: 2, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:24:14.230272   10206 sender.go:189] Sending Resources { request: 511166, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:24:39.327171   10206 sender.go:189] Sending Resources { request: 982312, add: 70, update: 1, delete: 0 edge add: 79 edge delete: 0 }
I0716 19:24:44.348251   10206 sender.go:189] Sending Resources { request: 285905, add: 158, update: 1, delete: 0 edge add: 364 edge delete: 0 }
I0716 19:24:49.372206   10206 sender.go:189] Sending Resources { request: 131778, add: 120, update: 17, delete: 0 edge add: 443 edge delete: 0 }
I0716 19:24:54.395627   10206 sender.go:189] Sending Resources { request: 887794, add: 13, update: 7, delete: 0 edge add: 48 edge delete: 0 }
I0716 19:24:59.420049   10206 sender.go:189] Sending Resources { request: 370505, add: 0, update: 22, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:25:04.443837   10206 sender.go:189] Sending Resources { request: 860222, add: 0, update: 4, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:25:14.491676   10206 sender.go:189] Sending Resources { request: 16738, add: 0, update: 9, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:25:19.513199   10206 sender.go:189] Sending Resources { request: 702617, add: 0, update: 9, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:25:24.539654   10206 sender.go:189] Sending Resources { request: 62388, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:25:29.563041   10206 sender.go:189] Sending Resources { request: 908784, add: 1, update: 1, delete: 0 edge add: 2 edge delete: 0 }
I0716 19:25:34.582925   10206 sender.go:189] Sending Resources { request: 817243, add: 1, update: 0, delete: 0 edge add: 2 edge delete: 0 }
I0716 19:25:44.636071   10206 sender.go:189] Sending Resources { request: 286188, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:25:49.660016   10206 sender.go:189] Sending Resources { request: 828925, add: 0, update: 2, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:26:04.732475   10206 sender.go:189] Sending Resources { request: 624454, add: 3, update: 0, delete: 3 edge add: 5 edge delete: 0 }
I0716 19:26:44.915831   10206 sender.go:189] Sending Resources { request: 134226, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:26:49.937883   10206 sender.go:189] Sending Resources { request: 438221, add: 0, update: 4, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:26:54.963239   10206 sender.go:189] Sending Resources { request: 268642, add: 1, update: 2, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:26:59.988547   10206 sender.go:189] Sending Resources { request: 258379, add: 1, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:05.011054   10206 sender.go:189] Sending Resources { request: 604549, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:10.039717   10206 sender.go:189] Sending Resources { request: 200486, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:20.086392   10206 sender.go:189] Sending Resources { request: 100983, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:25.111205   10206 sender.go:189] Sending Resources { request: 58175, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:35.151015   10206 sender.go:189] Sending Resources { request: 502075, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:40.179122   10206 sender.go:189] Sending Resources { request: 114238, add: 67, update: 7, delete: 2 edge add: 179 edge delete: 0 }
I0716 19:27:45.203612   10206 sender.go:189] Sending Resources { request: 289831, add: 6, update: 21, delete: 0 edge add: 47 edge delete: 0 }
I0716 19:27:50.228695   10206 sender.go:189] Sending Resources { request: 767214, add: 6, update: 2, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:27:55.254467   10206 sender.go:189] Sending Resources { request: 8563, add: 1, update: 5, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:00.276999   10206 sender.go:189] Sending Resources { request: 553365, add: 7, update: 4, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:05.304118   10206 sender.go:189] Sending Resources { request: 108384, add: 0, update: 3, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:10.331500   10206 sender.go:189] Sending Resources { request: 234007, add: 4, update: 7, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:15.352801   10206 sender.go:189] Sending Resources { request: 946694, add: 0, update: 4, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:20.378215   10206 sender.go:189] Sending Resources { request: 946759, add: 0, update: 3, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:25.400455   10206 sender.go:189] Sending Resources { request: 259807, add: 34, update: 3, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:30.436794   10206 sender.go:189] Sending Resources { request: 264699, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:45.515259   10206 sender.go:189] Sending Resources { request: 352377, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:28:50.540838   10206 sender.go:189] Sending Resources { request: 785994, add: 3, update: 0, delete: 1 edge add: 2 edge delete: 0 }
I0716 19:28:55.565636   10206 sender.go:189] Sending Resources { request: 994713, add: 0, update: 2, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:29:15.671840   10206 sender.go:189] Sending Resources { request: 342910, add: 0, update: 1, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:29:20.700940   10206 sender.go:189] Sending Resources { request: 199808, add: 0, update: 6, delete: 0 edge add: 0 edge delete: 0 }
I0716 19:29:25.727828   10206 sender.go:189] Sending Resources { request: 646337, add: 3, update: 5, delete: 0 edge add: 12 edge delete: 0 }
I0716 19:29:30.755547   10206 sender.go:189] Sending Resources { request: 729879, add: 0, update: 6, delete: 0 edge add: 0 edge delete: 0 }
