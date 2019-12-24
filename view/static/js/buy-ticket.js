{
    let view = {
        el: '#buy-ticket-box',
        template: `
        <form id="buy-ticket" class="buy-ticket" method="post">
            <div id="title" class="title">购买车票</div>
            <div class="input-box" id="input-box">
                <div>
                    <label for="">姓名</label>
                    <input type="text" autocomplete="off" name="name" value="">
                </div>
                <div>
                    <label for="">身份证号</label>
                    <input type="text" autocomplete="off" name="id_card" value="">
                </div>
                <div>
                    <label for="">手机号</label>
                    <input type="text" autocomplete="off" name="phone" value="">
                </div>
            </div>
            <div id="buy-success" class="buy-success">
                <svg t="1577109480241" class="successIcon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="951" width="120" height="120"><path d="M905.92 237.76a32 32 0 0 0-52.48 36.48A416 416 0 1 1 96 512a418.56 418.56 0 0 1 297.28-398.72 32 32 0 1 0-18.24-61.44A480 480 0 1 0 992 512a477.12 477.12 0 0 0-86.08-274.24z" fill="#9b2d7b" p-id="952"></path><path d="M630.72 113.28A413.76 413.76 0 0 1 768 185.28a32 32 0 0 0 39.68-50.24 476.8 476.8 0 0 0-160-83.2 32 32 0 0 0-18.24 61.44zM489.28 86.72a36.8 36.8 0 0 0 10.56 6.72 30.08 30.08 0 0 0 24.32 0 37.12 37.12 0 0 0 10.56-6.72A32 32 0 0 0 544 64a33.6 33.6 0 0 0-9.28-22.72A32 32 0 0 0 505.6 32a20.8 20.8 0 0 0-5.76 1.92 23.68 23.68 0 0 0-5.76 2.88l-4.8 3.84a32 32 0 0 0-6.72 10.56A32 32 0 0 0 480 64a32 32 0 0 0 2.56 12.16 37.12 37.12 0 0 0 6.72 10.56zM230.08 467.84a36.48 36.48 0 0 0 0 51.84L413.12 704a36.48 36.48 0 0 0 51.84 0l328.96-330.56A36.48 36.48 0 0 0 742.08 320l-303.36 303.36-156.8-155.52a36.8 36.8 0 0 0-51.84 0z" fill="#9b2d7b" p-id="953"></path></svg>
                <div class="text"></div>
            </div>
            <div id="buy-failed" class="buy-failed">
                <svg t="1577151989155" class="failedIcon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6777" width="120" height="120"><path d="M512.8 98.2C284.1 98.2 98 284.3 98 513.1c0 23.9 2.6 48.2 6.5 71 0 0.1 0.2 0.1 0.3 0.2 1.7 5.8 6.8 10.1 13.1 10.1 7.6 0 13.8-6.2 13.8-13.8 0-0.5-0.2-0.9-0.3-1.3l0.2-0.1c-3.7-21.5-6-43.5-6-66C125.6 299.4 299 126 512.8 126S900 299.2 900 513.1 726.7 900.3 512.8 900.3c-142.9 0-267.4-77.6-334.5-192.8 0-0.1-0.1-0.1-0.1-0.2l-0.6-0.9c-2.3-4.3-6.8-7.4-12-7.4-7.6 0-13.8 6.2-13.8 13.8 0 3 1.2 5.6 2.8 7.9l-0.2 0.1c71.9 123.6 205.5 207.1 358.5 207.1 228.7 0 414.8-186.1 414.8-414.8S741.6 98.2 512.8 98.2z" fill="#13227a" p-id="6778"></path><path d="M644.8 664.6L361.3 381.1c-5.4-5.4-5.4-14.2 0-19.6 5.4-5.4 14.2-5.4 19.6 0L664.4 645c5.4 5.4 5.4 14.2 0 19.6-5.4 5.4-14.2 5.4-19.6 0z" fill="#13227a" p-id="6779"></path><path d="M664.4 381.1L380.9 664.6c-5.4 5.4-14.2 5.4-19.6 0-5.4-5.4-5.4-14.2 0-19.6l283.5-283.5c5.4-5.4 14.2-5.4 19.6 0 5.4 5.5 5.4 14.2 0 19.6z" fill="#13227a" p-id="6780"></path></svg>
                <div class="text"></div>
            </div>
            <div class="change">
                <div class="cancel" id="cancel">取消</div>
                <input type="submit" class="buy" id="buy-button" value="购买">
            </div>
        </form>

        `,
        init() {
            this.$el = $(this.el);
        },
        render(data = {}) {
            // let str = ['name', 'id_card', 'phone'];
            // let html = this.template;
            // str.map((string)=>{
            //     html = html.replace(`__${string}__`, data[string] || '');
            // })
            this.$el.html(this.template);
        },
        show() {
            $('#ticket-list').css('filter', 'blur(6px)');
            this.$el.find('#buy-success').hide();
            this.$el.find('#buy-failed').hide();
            $('#unselect').show();
            this.$el.show();
            this.$el.find('[name="name"]').focus();
        },
        showSuccess(response) {
            this.$el.find('#buy-button').hide();
            this.$el.find('#input-box').hide();
            this.$el.find('#title').html('购买成功');
            this.$el.find('#cancel').html('确认');
            this.$el.find('#buy-success').show();
            this.$el.find('.text').html(`您的订单号为 ${response}<br />请妥善保管, 以便查询与退票`);
        },
        // showError() {
        //     this.$el.find('.buy').hide();
        //     this.$el.find('#input-box').hide();
        //     this.$el.find('#title').html('购买失败');
        //     this.$el.find('#cancel').html('确认');
        //     this.$el.find('#buy-failed').show();
        //     this.$el.find('.text').html(`购买车票失败`);
 
        // },
        hide() {
            $('#ticket-list').css('filter', 'blur(0px)');
            $('#unselect').hide();
            this.$el.hide();
            this.$el.find('#buy-button').show();
        }

    }

    let model = {
        data: {
            oticket_id: '',
            name: '',
            id_card: '',
            phone: ''
        },
        submitData() {
            let str = ['name', 'id_card', 'phone'];
            str.map((string)=>{
                this.data[string] = $(`input[name=${string}]`).val();
            })
            return $.post('/buyticket', this.data).then((response) => {
                return response;
            })
        }
    }

    let controller = {
        init(view, model) {
            this.view = view;
            this.model = model;
            this.view.init();
            this.view.hide();
            this.bindEvents();
            this.bindEventHubs();
        },
        bindEvents() {
            this.view.$el.on('submit', '#buy-ticket', (e)=>{
                e.preventDefault();
                this.model.submitData().then((response)=>{
                    this.view.showSuccess(response);
                })
            })

            this.view.$el.on('click', '#cancel', () => {
                this.view.hide();
            })
        },
        bindEventHubs() {
            window.eventHub.on('buyTicket',(ticketId)=>{
                this.view.render();
                this.model.data.oticket_id = ticketId;
                this.view.show();
            })
        }
    }
    controller.init(view, model);
}