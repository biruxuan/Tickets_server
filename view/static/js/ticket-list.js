{
    let view = {
        el: '#ticket-list-box',
        template: `
        <div class="ticket-list">
            <div class="title">车票购买系统</div>
            <ul>
            <li>
                <div class="ticket-id">班次</div><div class="departure-date">日期</div>
                <div class="departure-time">时间</div><div class="start-point">起点</div>
                <div class="end-point">终点</div><div class="travel-time">车程</div>
                <div class="rated-load">总票数</div><div class="booked-num">已售</div>
                <div class="buy">购买</div><div class="refund">退票</div>
            </ul>
            <div class="turing-page">
                <div class="prev" id="prev">上一页</div>
                <div class="pages" id="pages">__nowPage__ / __allPage__</div>
                <div class="next" id="next">下一页</div>
            </div>
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
                <div class="booked-num">${ticket.Booked_num}</div><div class="buy">购买</div>
                <div class="refund">退票</div>`).attr('ticket-id', ticket.Ticket_id);
                return $li;
            })
            liList.map((li, i) => {
                this.$el.find('ul').append(li);
            })
        },
        // show() {
        //     this.$el.css("z-index", '2');
        //     this.$el.fadeIn('fast');
        // },
        // hide() {
        //     this.$el.css("z-index", '-1');
        //     this.$el.fadeOut('fast');
        // },
        toPage(nowPage, allPage) {
            this.$el.find('#pages').html(`${nowPage} / ${allPage}`);
        }
    }

    let model = {
        data: {
            tickets: [],
            nowPage: 1,
            allPage: 1
        },
        getTickets() {
            return $.get('/alltickets').then((response)=> {
                this.data.tickets = response;
                this.data.allPage = parseInt((this.data.tickets.length - 1) / 8) + 1;
                if (this.data.nowPage > this.data.allPage) {
                    this.data.nowPage = this.data.allPage;
                }
                return response;
            })
            // let ticketsRequest = new XMLHttpRequest();
            // ticketsRequest.open("GET", "/alltickets");
            // ticketsRequest.send();
            // ticketsRequest.onreadystatechange = ()=> {
            //         if (ticketsRequest.readyState === 4 && ticketsRequest.status === 200) {
            //             this.data.tickets = ticketsRequest.responseText;
            //             console.log(this.data);
            //             this.data.allPage = parseInt((this.data.tickets.length - 1) / 8) + 1;
            //             if (this.data.nowPage > this.data.allPage) {
            //                 this.data.nowPage = this.data.allPage;
            //             }
            //             return true;
            //         } else if (ticketsRequest.readyState === 4 && ticketsRequest.status === 404) {
            //             console.log('读取失败');
            //         }
            //     }
                // this.data.tickets = testData;

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
            // this.view.render({});
            this.model.getTickets().then(()=>{
                   this.view.render(this.model.data);
                })

            this.bindEvents();
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
        bindEvents() {
            this.view.$el.on('click', '#next', () => { //跳转到下一页
                this.toPage('next');
            })

            this.view.$el.on('click', '#prev', () => { //跳转到上一页
                this.toPage('prev');
            })

            this.view.$el.on('hover', '#li', (e) => {

            })
        },
    }
    controller.init(view, model);
}