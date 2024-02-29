// Copyright 2020 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httputil

// BasicAuth is used for authorization on /metrics handlers
type BasicAuth struct {
	walletAddr string `yaml:"walletAddr"`
}

/*
*
登录验证
MakeUserAuthAPI turns a pkg.JSONRequestHandler function into an http.Handler which authenticates the request.
*/
//func MakeUserAuthAPI(
//	metricsName string,
//	db *storage.Database,
//	f func(*http.Request, *storage.User) pkg.JSONResponse,
//) http.Handler {
//	h := func(req *http.Request) pkg.JSONResponse {
//		account, jsonRes := auth.VerifyUserFromRequest(req, db)
//		if jsonRes != nil {
//			return *jsonRes
//		}
//		// add the user ID to the logger
//		logger := pkg.GetLogger((req.Context()))
//		//logger = logger.WithField("user_id", device.UserID)
//		req = req.WithContext(pkg.ContextWithLogger(req.Context(), logger))
//		ret := f(req, account)
//		return ret
//	}
//	return MakeExternalAPI(metricsName, h)
//}

/*
*
免登录验证
MakeExternalAPI turns a pkg.JSONRequestHandler function into an http.Handler.
This is used for APIs that are called from the internet.
*/
//func MakeExternalAPI(metricsName string, f func(*http.Request) pkg.JSONResponse) http.Handler {
//
//	h := pkg.MakeJSONAPI(pkg.NewJSONRequestHandler(f))
//	corfunc := func(w http.ResponseWriter, req *http.Request) {
//		nextWriter := w
//		h.ServeHTTP(nextWriter, req)
//	}
//	httpHandler := http.HandlerFunc(corfunc)
//
//	return httpHandler
//}
