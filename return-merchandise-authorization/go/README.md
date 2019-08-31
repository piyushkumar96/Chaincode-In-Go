# Commands for install, instantiate, invoke, query with chaincode

#create package of chaindode

docker exec -it cli peer chaincode package -n rma-test  -p github.com/chaincode/chaincode_example02/go  -v 1.0 -s -S -i "AND('OrgA.admin')" ccpack.out

#sign the package

docker exec -it cli peer chaincode signpackage ccpack.out signedccpack.out

#Install chaincode

docker exec -it cli peer chaincode install -n rma-test  -v 1.0 -p github.com/chaincode/chaincode_example02/go

#Instaintiating chaincode

docker exec -it cli peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test   -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"

#Invoke Chaincode

docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test -c '{"Args":[ "createTicket","1"]}'

docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test -c '{"Args":[ "updates3b11","1","Some changes", "UUUUUU"]}'

docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test -c '{"Args":[ "updates3b13","1","Some changes", "UUUUUU"]}'

 docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test -c '{"Args":[ "updates3b3eta","1","Some changes", "UUUUUU"]}'


docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test -c '{"Args":[ "updates3b3pod","1","Some changes", "UUUUUU"]}'

docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n rma-test -c '{"Args":[ "updates3b3sir","1","Some changes", "UUUUUU"]}'


#Query Chaincode

docker exec -it cli peer chaincode query -o orderer.example.com:7050  -C mychannel -n rma-test -c '{"Args":[ "getAllMessagesById","1"]}'

docker exec -it cli peer chaincode query -o orderer.example.com:7050  -C mychannel -n rma-test -c '{"Args":[ "getAllMessages",""]}'




