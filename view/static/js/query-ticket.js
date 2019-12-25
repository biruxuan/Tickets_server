{
    let view = {
        el: '#query-ticket-box',
        template: `
        <form id="query-ticket" class="query-ticket" method="post">
            <div id="title" class="title">查询订单</div>
            <div class="input-box" id="input-box">
                <div>
                    <label for="">订单号</label>
                    <input type="text" autocomplete="off" name="order_id" value="" placeholder="必填">
                </div>
                <div>
                    <label for="">身份证号</label>
                    <input type="text" autocomplete="off" name="phone" value="" placeholder="选填">
                </div>
            </div>
            <div id="query-success" class="query-success">
                <div class="point">
                    <div class="start-point">__Start_point__</div>
                    <div class="arrive">
                        <div class="train-id">__TrainID__</div>
                        <svg t="1577174797239" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2742" width="40" height="40"><path d="M121.626556 637.344398h775.643884a39.897759 39.897759 0 0 0 39.226423-32.632033c3.777328-20.395021-2.851054-41.333909-17.675685-55.839867l-253.408132-247.926971a31.25112 31.25112 0 0 0-45.748581 2.188216 35.844249 35.844249 0 0 0 2.608863 48.994788L851.917012 572.080332H125.25517a34.548315 34.548315 0 0 0-34.497328 32.632033 30.915452 30.915452 0 0 0 30.868714 32.632033z" p-id="2743" fill="#9b2d7b"></path></svg>
                    </div>
                    <div class="end-point">__End_point__</div>
                </div>
                <div class="departure-time">发车时间：__Departure_date__ __Departure_time__</div>
                <div class="travel-time">行程时间：__TravelTime__ 小时</div>
                <div class="name">车票姓名：__Name__</div>
                <div class="id-card">身份证号：__IDCard__</div>
            </div>
            <div id="refund-success" class="refund-success">
                <svg t="1577109480241" class="successIcon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="951" width="120" height="120"><path d="M905.92 237.76a32 32 0 0 0-52.48 36.48A416 416 0 1 1 96 512a418.56 418.56 0 0 1 297.28-398.72 32 32 0 1 0-18.24-61.44A480 480 0 1 0 992 512a477.12 477.12 0 0 0-86.08-274.24z" fill="#9b2d7b" p-id="952"></path><path d="M630.72 113.28A413.76 413.76 0 0 1 768 185.28a32 32 0 0 0 39.68-50.24 476.8 476.8 0 0 0-160-83.2 32 32 0 0 0-18.24 61.44zM489.28 86.72a36.8 36.8 0 0 0 10.56 6.72 30.08 30.08 0 0 0 24.32 0 37.12 37.12 0 0 0 10.56-6.72A32 32 0 0 0 544 64a33.6 33.6 0 0 0-9.28-22.72A32 32 0 0 0 505.6 32a20.8 20.8 0 0 0-5.76 1.92 23.68 23.68 0 0 0-5.76 2.88l-4.8 3.84a32 32 0 0 0-6.72 10.56A32 32 0 0 0 480 64a32 32 0 0 0 2.56 12.16 37.12 37.12 0 0 0 6.72 10.56zM230.08 467.84a36.48 36.48 0 0 0 0 51.84L413.12 704a36.48 36.48 0 0 0 51.84 0l328.96-330.56A36.48 36.48 0 0 0 742.08 320l-303.36 303.36-156.8-155.52a36.8 36.8 0 0 0-51.84 0z" fill="#9b2d7b" p-id="953"></path></svg>
                <div class="text"></div>
            </div>
            <div class="change">
                <div class="cancel" id="cancel">取消</div>
                <div class="delete" id="delete">退票</div>
                <input type="submit" class="query-button" id="query-button" value="查询">
            </div>
        </form>
        `,
        init() {
            this.$el = $(this.el);
        },
        render() {
            this.$el.html(this.template);
        },
        show() {
            $('#ticket-list').css('filter', 'blur(6px)');
            this.$el.find('#query-success').hide();
            this.$el.find('#refund-success').hide();
            $('#unselect').show();
            this.$el.show();
            this.$el.find('[name="order_id"]').focus();
        },
        renderOrder(data = {}) {
            let str = ['TrainID', 'Departure_date', 'Departure_time', 'Start_point', 'End_point', 'TravelTime', 'Name', 'IDCard'];
            let html = this.template;
            str.map((string) => {
                html = html.replace(`__${string}__`, data[string] || '');
            })
            this.$el.html(html);
        },
        showOrder() {
            // this.$el.show();
            this.$el.find('#refund-success').hide();
            this.$el.find('#query-button').hide();
            this.$el.find('#input-box').hide();
            this.$el.find('#delete').show();
            this.$el.find('#title').html('订单信息');
            this.$el.find('#cancel').html('确认');
            this.$el.find('#query-success').show();
        },
        refundSuccess(data) {
            this.$el.find('#query-success').hide();
            this.$el.find('#delete').hide();
            this.$el.find('#title').html('退票成功');
            this.$el.find('.text').html(`恭喜! 订单 ${data.order_id} 已成功退票`);
            this.$el.find('#refund-success').show();
        },
        hide() {
            $('#ticket-list').css('filter', 'blur(0px)');
            $('#unselect').hide();
            this.$el.find('#query-button').show();
            this.$el.hide();
        }
    }

    let model = {
        data: {
            query: {
                order_id: ''
            },
            order: []
        },
        queryOrder() {
            let str = ['order_id', 'phone', 'id_card'];
            this.data.query.order_id = $(`input[name="order_id"]`).val();
            return $.post('/queryticket', this.data.query).then((response) => {
                this.data.order = response;
                return response;
            })
        },
        refundOrder() {
            return $.post('/refundorder', 
                {
                    order_id: this.data.query.order_id,
                    ticket_id: this.data.order.TicketID
                }).then((response) => {
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
            this.view.$el.on('submit', '#query-ticket', (e) => {
                e.preventDefault();
                this.model.queryOrder().then((response) => {
                    this.view.renderOrder(response);
                    this.view.showOrder();
                })
            })

            this.view.$el.on('click', '#cancel', () => {
                this.view.hide();
                window.eventHub.emit('renderView');
            })

            this.view.$el.on('click', '#delete', () => {
                this.model.refundOrder().then(() => {
                    this.view.refundSuccess(this.model.data.query);
                })
            })
        },
        bindEventHubs() {
            window.eventHub.on('queryOrder', (data) => {
                this.view.render();
                this.view.show();
            })
        }
    }

    controller.init(view, model);
}