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
                    <div class="start-point">青岛</div>
                    <div class="arrive">
                        <div class="train-id">Z101</div>
                        <svg t="1577174797239" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2742" width="40" height="40"><path d="M121.626556 637.344398h775.643884a39.897759 39.897759 0 0 0 39.226423-32.632033c3.777328-20.395021-2.851054-41.333909-17.675685-55.839867l-253.408132-247.926971a31.25112 31.25112 0 0 0-45.748581 2.188216 35.844249 35.844249 0 0 0 2.608863 48.994788L851.917012 572.080332H125.25517a34.548315 34.548315 0 0 0-34.497328 32.632033 30.915452 30.915452 0 0 0 30.868714 32.632033z" p-id="2743" fill="#9b2d7b"></path></svg>
                    </div>
                    <div class="end-point">北京</div>
                </div>
                <div class="departure-time">发车时间：2019-12-24 16:00:00</div>
                <div class="name">购票人：郭宇航</div>
            </div>
            <div id="query-failed" class="query-failed">
                <svg t="1577151989155" class="failedIcon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6777" width="120" height="120"><path d="M512.8 98.2C284.1 98.2 98 284.3 98 513.1c0 23.9 2.6 48.2 6.5 71 0 0.1 0.2 0.1 0.3 0.2 1.7 5.8 6.8 10.1 13.1 10.1 7.6 0 13.8-6.2 13.8-13.8 0-0.5-0.2-0.9-0.3-1.3l0.2-0.1c-3.7-21.5-6-43.5-6-66C125.6 299.4 299 126 512.8 126S900 299.2 900 513.1 726.7 900.3 512.8 900.3c-142.9 0-267.4-77.6-334.5-192.8 0-0.1-0.1-0.1-0.1-0.2l-0.6-0.9c-2.3-4.3-6.8-7.4-12-7.4-7.6 0-13.8 6.2-13.8 13.8 0 3 1.2 5.6 2.8 7.9l-0.2 0.1c71.9 123.6 205.5 207.1 358.5 207.1 228.7 0 414.8-186.1 414.8-414.8S741.6 98.2 512.8 98.2z" fill="#13227a" p-id="6778"></path><path d="M644.8 664.6L361.3 381.1c-5.4-5.4-5.4-14.2 0-19.6 5.4-5.4 14.2-5.4 19.6 0L664.4 645c5.4 5.4 5.4 14.2 0 19.6-5.4 5.4-14.2 5.4-19.6 0z" fill="#13227a" p-id="6779"></path><path d="M664.4 381.1L380.9 664.6c-5.4 5.4-14.2 5.4-19.6 0-5.4-5.4-5.4-14.2 0-19.6l283.5-283.5c5.4-5.4 14.2-5.4 19.6 0 5.4 5.5 5.4 14.2 0 19.6z" fill="#13227a" p-id="6780"></path></svg>
                <div class="text"></div>
            </div>
            <div class="change">
                <div class="cancel" id="cancel">取消</div>
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
            this.$el.find('#query-failed').hide();
            $('#unselect').show();
            this.$el.show();
            this.$el.find('[name="order_id"]').focus();
        },
        renderOrder(data = {}) {
            let str = [''];
            let html = this.template;
            str.map((string) => {
                html = html.replace(`__${string}__`, data[string] || '');
            })
            this.$el.html(html);
        },
        showOrder() {
            this.$el.show();
            this.$el.find('#query-failed').hide();
            this.$el.find('#query-button').hide();
            this.$el.find('#input-box').hide();
            this.$el.find('#title').html('订单信息');
            this.$el.find('#cancel').html('确认');
            this.$el.find('#query-success').show();
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
                console.log(response);
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
                    // this.view.renderOrder(response);
                    this.view.showOrder();
                })
            })

            this.view.$el.on('click', '#cancel', () => {
                this.view.hide();
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