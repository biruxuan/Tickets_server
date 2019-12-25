{
    let view = {
        el: '#ticket-list',
        template: `
            <div class="title">车票购买系统</div>
            <div id="now-time" class="now-time"></div>
            <div id="query-order" class="query-order">查询订单</div>
            <ul>
            <li>
                <div class="train-id">班次</div><div class="departure-date">日期</div>
                <div class="departure-time">时间</div><div class="start-point">起点</div>
                <div class="end-point">终点</div><div class="travel-time">车程</div>
                <div class="rated-load">总票数</div><div class="booked-num">已售</div>
                <div class="buy"></div>
            </ul>
            <div class="turing-page">
                <div class="prev" id="prev">上一页</div>
                <div class="pages" id="pages">__nowPage__ / __allPage__</div>
                <div class="next" id="next">下一页</div>
            </div>
        `,
        init() {
            this.$el = $(this.el);
        },
        render({ tickets = [], nowPage, allPage, time }) {    //渲染车票数据
            let html = this.template;
            //替换页数并渲染
            html = html.replace('__nowPage__', nowPage).replace('__allPage__', allPage);
            this.$el.html(html);

             //获取当前时间的毫秒
            let nDSeconds = new Date().getTime();
            //开始循环渲染每条车票数据
            let liList = tickets.map((ticket) => {
                //添加一条车票数据li元素
                let $li = $('<li></li>').html(`<div class="train-id">${ticket.Train_id}</div>
                            <div class="departure-date">${ticket.Departure_date}</div><div class="departure-time">
                            ${ticket.Departure_time}</div><div class="start-point">${ticket.Start_point}</div>
                            <div class="end-point">${ticket.End_point}</div><div class="travel-time">
                            ${ticket.Travel_time}</div><div class="rated-load">${ticket.Rated_load}</div>
                            <div class="booked-num">${ticket.Booked_num}</div><div id="buy" class="buy">购买</div>`)
                            .attr('ticket-id', ticket.Ticket_id);
                //获取车票发车时间的毫秒
                let tDate = `${ticket.Departure_date.replace(new RegExp("-","gm"),"/")} ${ticket.Departure_time}`;
                let tDSeconds = new Date(tDate).getTime();
                //判断已发车 和 已售光
                if (nDSeconds >= tDSeconds) {
                    $li.find('#buy').addClass('time-out').text('已发车');
                } else if(ticket.Booked_num >= ticket.Rated_load) {
                    $li.find('#buy').addClass('time-out').text('已售光');
                } else {
                    $li.find('#buy').addClass('time');
                }
                
                return $li;
            })
            liList.map((li, i) => {     //将车票数据添加到ul元素内
                this.$el.find('ul').append(li);
            })
        },
        toPage(nowPage, allPage) {  //渲染页数
            this.$el.find('#pages').html(`${nowPage} / ${allPage}`);
        },
        setNowTime(time) {  //渲染时间
            this.$el.find('#now-time').html(time);
        }
    }

    let model = {
        data: {
            tickets: [],
            nowPage: 1,
            allPage: 1,
            time: ''
        },
        getTickets() {  //从后端读取车票数据
            return $.get('/allTickets').then((response)=> {
                this.data.tickets = response;
                //计算显示总页数
                this.data.allPage = parseInt((this.data.tickets.length - 1) / 8) + 1;
                if (this.data.nowPage > this.data.allPage) {
                    this.data.nowPage = this.data.allPage;
                }
                return response;
            })

        },
        getNowTime() {  //获取当前时间
            let nowTime = new Date();
            let hour = nowTime.getHours() >= 10 ? nowTime.getHours() : '0' + nowTime.getHours();
            let minute = nowTime.getMinutes() >= 10 ? nowTime.getMinutes() : '0' + nowTime.getMinutes();
            let second = nowTime.getSeconds() >= 10 ? nowTime.getSeconds() : '0' + nowTime.getSeconds();
            this.data.time = `${nowTime.getFullYear()}-${nowTime.getMonth()+1}-${nowTime.getDate()} ${hour}:${minute}:${second}`;
        },
        toPage(allLi, pagedNum) {   //跳转到pagedNum
            let showLi = 8;
            let allPageNum = this.data.allPage;
            if (1 <= pagedNum && pagedNum <= allPageNum) {
                for (let i = 1; i <= allLi.length - 1; i++) {
                    if (i <= (pagedNum - 1) * showLi) {
                        $(allLi[i]).hide();
                    } else {
                        $(allLi[i]).show();
                    }
                }
                this.data.nowPage = pagedNum;
            }
        }
    }
let controller = {
        init(view, model) {
            this.view = view;
            this.model = model;
            this.view.init();
            this.view.render({});   
            this.renderView();
            this.getRenderSeconds();
            this.bindEvents();
            this.bindEventHubs();
            this.getNowTime();
        },
        toPage(status) {
            let nowPage = this.model.data.nowPage;
            let allPage = this.model.data.allPage;
            let allLi = this.view.$el.find('li');
            let addPageNum = 0;

            switch (status) {
                case 'prev':
                    addPageNum = -1;
                    break;
                case 'now':
                    addPageNum = 0;
                    break;
                case 'next':
                    addPageNum = 1;
                    break;
                case 'last':
                    addPageNum = allPage - nowPage;
                    break;
            }

            this.model.toPage(allLi, nowPage + addPageNum);
            this.view.toPage(this.model.data.nowPage, allPage);
        },
        getNowTime() {  //循环获取当前时间  间隔 500ms
            this.model.getNowTime();
            this.view.setNowTime(this.model.data.time);
            setInterval(() => {
                this.model.getNowTime();
                this.view.setNowTime(this.model.data.time);
            }, 500);
        },
        bindEvents() {
            this.view.$el.on('click', '#next', () => { //跳转到下一页
                this.toPage('next');
            })

            this.view.$el.on('click', '#prev', () => { //跳转到上一页
                this.toPage('prev');
            })

            this.view.$el.on('click', '.time', (e) => {  //购买车票按钮
                let ticketId = $(e.currentTarget).parent().attr('ticket-id');
                window.eventHub.emit('buyTicket', ticketId);
            })

            this.view.$el.on('click', '#query-order', (e) => {  //发布查询订单
                window.eventHub.emit('queryOrder');
            })
        },
        bindEventHubs() {
            window.eventHub.on('renderView', () => {
                this.renderView();
            })
        },
        getRenderSeconds() {    //当秒数为0时调用渲染车票函数   判断间隔 1s
            let second;
             setInterval(()=>{
                second = new Date().getSeconds();
                if (second === 0) {
                    this.renderView();
                }
             }, 1000)
        },
        renderView() {  //渲染车票
            this.model.getTickets().then(() => {
                this.view.render(this.model.data);
             });
        }
    }
    controller.init(view, model);
}