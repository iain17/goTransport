module goTransport {

    import IPromise = goTransport.IPromise;
    export abstract class Client {

        protected static instance: Client;
        private messageManager = new MessageManager();

        constructor() {
            Client.instance = this;
        }

        public connect(url : string): IPromise<{}> {
            return this.messageManager.connect(url);
        }

        public method(name: string, parameters: any[], timeout: number = 3000): IPromise<{}> {
            let message = new MessageMethod(name, parameters);
            this.messageManager.send(message);
            var promise = message.getPromise();
            promise.setTimeOut(timeout);
            return promise.promise;
        }

        public onConnect(): IPromise<{}> {
            return this.messageManager.getConnectedPromise();
        }

    }

}