/*
 * @Description: Do not edit
 * @Author: Jianxuesong
 * @Date: 2021-06-14 10:50:40
 * @LastEditors: Jianxuesong
 * @LastEditTime: 2021-06-16 11:38:53
 * @FilePath: /Melon/middleware/Gray.go
 */
package middleware

// func Gray() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// fmt.Println(c.GetHeader("X-Real-IP"))
// 		isGray := true
// 		url := "http://127.0.0.1:8001/v1/geo"
// 		params := map[string]string{
// 			// 124.236.219.174 石家庄 36.110.66.197 北京
// 			"ip":     c.ClientIP(), // c.GetHeader("X-Real-IP"),
// 			"lon":    c.Query("lon"),
// 			"lat":    c.Query("lat"),
// 			"_debug": c.Query("_debug"),
// 			"_flush": c.Query("_flush"),
// 		}

// 		if params["lon"] == "" {
// 			params["lon"] = c.GetHeader("lon")
// 		}

// 		if params["lat"] == "" {
// 			params["lat"] = c.GetHeader("lat")
// 		}

// 		// fmt.Println(params)

// 		rm := model.NewRegionModel(c)

// 		startTime := time.Now()
// 		tmp, err := proxy.SimpleClient(url, "GET", nil, params)
// 		rm.SetDebug(1, "{Get geo info, url: %s?%s (%s)}", url, public.UrlEncode(params), public.TimeCost(startTime))
// 		if err != nil {
// 			logzap.Ex(c.Request.Context(), "mango", "Middleware Geo proxy error: %s", err.Error())
// 		} else {
// 			if tmp.StatusCode != 200 {
// 				logzap.Ex(c, "mango", "Middleware Geo http error: %d", tmp.StatusCode)
// 			} else {
// 				if code := jsoniter.Get(tmp.Body, "code").ToInt(); code != 1 {
// 					logzap.Ex(c, "mango", "Middleware Geo error: %s", string(tmp.Body))
// 				} else {
// 					location := jsoniter.Get(tmp.Body, "data", "location").ToString()
// 					c.Set("location", location)
// 					// fmt.Println(location)
// 					// rm := model.NewRegionModel(c)
// 					isGray = rm.IsGray(location)
// 				}
// 			}
// 		}

// 		// uid, _ := strconv.Atoi(c.GetHeader("uid")) // header中获取uid
// 		token := c.GetHeader("User-Token") // header中获取token
// 		// token = "5b7aab734acfc49251527d7b16729084f0fc4f64e169043832a6e70eeffb843"
// 		user, err := service.UserTokenS(c).GetUserByToken(token)
// 		if err != nil {
// 			logzap.Ex(c, "mango", "Get user by token %s", err.Error())
// 		}

// 		// fmt.Println(isGray, user, "token:", token)
// 		// is_gray = true
// 		if isGray && user != nil {
// 			whiteUser, err := service.UserS(c).GetWhiteUser(false)
// 			// fmt.Println(whiteUser)
// 			if err != nil {
// 				logzap.Ex(c, "mango", "Middleware Gray", err.Error())
// 			} else {
// 				if _, inArray := whiteUser[user.UserID]; inArray {
// 					rm.SetDebug(1, "{The user is customer white, uid: %d}", user.UserID)
// 					isGray = false
// 				}
// 			}
// 		}

// 		c.Set("uid", 0)
// 		if user != nil {
// 			c.Set("uid", user.UserID)
// 		}

// 		// fmt.Println("is_gray", is_gray)
// 		// c.Set("uid", 15)
// 		c.Set("is_gray", isGray)

// 		// defer func() {
// 		// 	if err := recover(); err != nil {
// 		// 		logcus.Error(fmt.Sprintf("Cos Panic info is: %v", err))
// 		// 	}
// 		// }()

// 		c.Next()
// 	}
// }
