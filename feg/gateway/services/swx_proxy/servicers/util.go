/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package servicers

import (
	"github.com/fiorix/go-diameter/v4/diam"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *swxProxy) sendDiameterMsg(msg *diam.Message, retryCount uint) error {
	conn, err := s.connMan.GetConnection(s.smClient, s.config.ServerCfg)
	if err != nil {
		return err
	}
	err = conn.SendRequest(msg, retryCount)
	if err != nil {
		err = status.Errorf(codes.DataLoss, err.Error())
	}
	return err
}

func (s *swxProxy) IsHlrClient(imsi string) bool {
	if s != nil {
		return s.config.IsHlrClient(imsi)
	}
	return false
}