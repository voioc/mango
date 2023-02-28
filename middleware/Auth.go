package middleware

// func Auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := strings.Replace(c.GetHeader("Authorization"), " ", "", -1)
// 		tokenClaims, err := jwt.ParseWithClaims(token, &define.Claims{}, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(common.JWT_SECRET), nil
// 		})

// 		if err != nil {
// 			logzap.Ex(c, "mango", "jwt vaild error: %s | token: %s", err.Error(), token)
// 			c.JSON(http.StatusOK, common.Error(c, common.ERROR_AUTH))
// 			c.Abort()
// 			return
// 		}

// 		claims, ok := tokenClaims.Claims.(*define.Claims)
// 		if !ok || !tokenClaims.Valid {
// 			logzap.Ex(c, "mango", "jwt vaild error: %s", token)
// 			c.JSON(http.StatusOK, common.Error(c, common.ERROR_AUTH))
// 			c.Abort()
// 			return
// 		}

// 		// fmt.Println(claims)
// 		c.Set("uid", claims.UID)
// 		c.Next()
// 	}
// }

// func AuthTK() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		time := c.GetHeader("t")
// 		uri := c.Request.URL.Path
// 		param := decodeParam(c.Request.URL.RawQuery)
// 		secret := "XSpeUFjJ"
// 		str := fmt.Sprintf("%s%s%s%s", uri, param, time, secret)

// 		data := []byte(str)
// 		has := md5.Sum(data)
// 		token := fmt.Sprintf("%x", has)

// 		tk := c.GetHeader("TK")
// 		if tk != token {
// 			str := fmt.Sprintf("%s|%s|%s|%s", uri, param, time, secret)
// 			public.FilePutContents("/tmp/token_go.log", fmt.Sprintf("ip: %s | tk: %s | str: %s", c.ClientIP(), tk, str), true)
// 		}

// 		c.Next()
// 	}
// }

// func decodeParam(p string) string {
// 	pArray := strings.Split(p, "&") //1、先按&切割
// 	keys := []string{}
// 	data := make(map[string]string)
// 	for _, v := range pArray {
// 		// 2、按照 = 切分组装map
// 		vs := strings.Split(v, "=")
// 		if len(vs) == 2 && vs[0] != "_debug" && vs[0] != "_flush" {
// 			value, err := url.QueryUnescape(vs[1]) // 从上面打印的字符，可以看出被urlescape过，因此要Unescape
// 			if err != nil {
// 				logzap.Ex(context.Background(), "tk error", "decode param: %s | key: %s | param: %s", err.Error(), vs[0], p)
// 				return ""
// 			}
// 			data[vs[0]] = value
// 			keys = append(keys, vs[0])
// 		}
// 	}

// 	sort.Strings(keys)
// 	s := ""
// 	for _, key := range keys {
// 		s = fmt.Sprintf("%s%v", s, data[key])
// 	}

// 	// s = s[:len(s)-1]

// 	return s
// }
