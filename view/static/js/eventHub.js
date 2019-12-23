window.eventHub = {
    events: {},
    emit(eventName, data) {
        let fnList = this.events[eventName];
        fnList.map((fn) => {
            fn.call(undefined, data);
        })
    },
    on(eventName, fn) {
        if (this.events[eventName] === undefined) {
            this.events[eventName] = [];
        }
        this.events[eventName].push(fn);
    }
}