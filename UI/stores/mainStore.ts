import { defineStore } from "pinia"

const url: string = "localhost:3001"
var client:WebSocket


export const useMainStore = defineStore('mainStore', {
    state() {
        return {
            character: "Pusheen",
            messages: [],
            status:"Not Connected", 
        }
    },
    getters: {
        getCharacter(): string {
            return this.character
        },
        getMessages(): Array<string> {
            return this.messages
        },
        getStatus():string{
            return this.status
        }
    },
    actions: {
        setChar(name: string) {
            this.character = name
        },
        async connect() {
            try{
                if(client){
                    client.close()
                }
                client =  await new WebSocket(`ws://${url}/ws`)
                client.addEventListener("open",()=>{
                    this.setStatus("Connected")
                })
                client.addEventListener("message",(msg)=>{
                    this.addMessage(msg)
                })
                client.addEventListener("close", ()=>{
                    this.setStatus("Not Connected")
                })
            }
            catch(err){
                console.log(err)
            }
        },
        setStatus(status:string){
            this.status = status
        },
        sendMessage(eventName:string,payload:string){
            let newPayload = {
                eventName,
                payload,
                user: this.character
            }
            console.log(newPayload)
            client.send(JSON.stringify(newPayload))
        },
        addMessage(msg:MessageEvent){
            let payload = JSON.parse(msg.data)
            payload.user === this.character ? payload.ownMessage = true : payload.ownMessage = false
            this.messages.push(payload)
        }
    }

})