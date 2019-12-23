{
    let view = {
        el: '#ticket-list',
        template: `
            <div class="title">车票购买系统</div>
            <div id="now-time" class="now-time"></div>
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
        render({ tickets = [], nowPage, allPage }) {
            let html = this.template;
            html = html.replace('__nowPage__', nowPage).replace('__allPage__', allPage);
            this.$el.html(html);
            let number = 1;
            let liList = tickets.map((ticket) => {
                let $li = $('<li></li>').html(`<div class="train-id">${ticket.Train_id}</div>
                <div class="departure-date">${ticket.Departure_date}</div><div class="departure-time">
                ${ticket.Departure_time}</div><div class="start-point">${ticket.Start_point}</div>
                <div class="end-point">${ticket.End_point}</div><div class="travel-time">
                ${ticket.Travel_time}</div><div class="rated-load">${ticket.Rated_load}</div>
                <div class="booked-num">${ticket.Booked_num}</div><div id="buy" class="buy">购买</div>`)
                .attr('ticket-id', ticket.Ticket_id);
                return $li;
            })
            liList.map((li, i) => {
                this.$el.find('ul').append(li);
            })
        },
        toPage(nowPage, allPage) {
            this.$el.find('#pages').html(`${nowPage} / ${allPage}`);
        },
        setNowTime(time) {
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
        getTickets() {
            return $.get('/allTickets').then((response)=> {
                this.data.tickets = response;
                this.data.allPage = parseInt((this.data.tickets.length - 1) / 8) + 1;
                if (this.data.nowPage > this.data.allPage) {
                    this.data.nowPage = this.data.allPage;
                }
                return response;
            })

        },
        getNowTime() {
            let nowTime = new Date();
            let hour = nowTime.getHours() > 9 ? nowTime.getHours() : '0' + nowTime.getHours();
            let minute = nowTime.getMinutes() > 9 ? nowTime.getMinutes() : '0' + nowTime.getMinutes();
            let second = nowTime.getSeconds() > 9 ? nowTime.getSeconds() : '0' + nowTime.getSeconds();
            this.data.time = `${nowTime.getFullYear()}-${nowTime.getMonth()+1}-${nowTime.getDate()} ${hour}:${minute}:${second}`;
        },
        toPage(allLi, pagedNum) {
            let allPageNum = this.data.allPage;
            if (1 <= pagedNum && pagedNum <= allPageNum) {
                for (let i = 1; i <= allLi.length - 1; i++) {
                    if (i <= (pagedNum - 1) * 8) {
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
            this.model.getTickets().then(() => {
                   this.view.render(this.model.data);
                });
            this.bindEvents();
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
        getNowTime() {
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

            this.view.$el.on('click', '#buy', (e) => {
                let ticketId = $(e.currentTarget).parent().attr('ticket-id');
                window.eventHub.emit('buyTicket', ticketId);
            })
        },
    }
    controller.init(view, model);
}