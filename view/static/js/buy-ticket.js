{
    let view = {
        el: '#buy-ticket-box',
        template: `
        <form class="buy-ticket" method="post">
        <div class="title">购买车票</div>
        <div class="input-box">
            <div>
                <label for="">姓名</label>
                <input type="text" autocomplete="off" name="name" value="__name__">
            </div>
            <div>
                <label for="">身份证号</label>
                <input type="text" autocomplete="off" name="id_card" value="__id_card__">
            </div>
            <div>
                <label for="">手机号</label>
                <input type="text" autocomplete="off" name="phone" value="__phone__">
            </div>
        </div>
        <div class="change">
            <div class="cancel" id="cancel">取消</div>
            <input type="submit" class="buy" value="购买">
        </div>
    </form>
        `,
        init() {
            this.$el = $(this.el);
        },
        render(data = {}) {
            let str = ['name', 'id_card', 'phone'];
            let html = this.template;
            str.map((string)=>{
                html = html.replace(`__${string}__`, data[string] || '');
            })
            this.$el.html(html);
        },
        show() {
            $('#ticket-list').css('filter', 'blur(6px)');
            $('#unselect').show();
            this.$el.show();
        },
        hide() {
            $('#ticket-list').css('filter', 'blur(0px)');
            $('#unselect').hide();
            this.$el.hide();
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
            return $.post('/buyticket', this.data).then(() => {
                return true;
            })
        },
        
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
            this.view.$el.on('submit', 'form', (e)=>{
                e.preventDefault();
                this.model.submitData().then(()=>{
                    this.view.hide();
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