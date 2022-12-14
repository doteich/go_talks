import { defineStore } from "pinia"

export const useMainStore = defineStore('mainStore', {
    state() {
        return {
            character: "Pusheen"
        }
    },
    getters: {
        getCharacter(): string {
            return this.character
        }
    },
    actions: {
        setChar(name: string) {
            this.character = name
        }
    }

})