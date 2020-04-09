// Tank define tank 
function Tank(x, y, direct, speed, type) {
    this.x = x;
    this.y = y;
    this.direct = direct;
    this.speed = speed;
    this.isLive = true;
    this.type = type;
}

Tank.prototype.moveUp = function () {
    this.direct = 0;
    this.y -= this.speed;
}

Tank.prototype.moveRight = function () {
    this.direct = 1;
    this.x += this.speed;
}

Tank.prototype.moveDown = function () {
    this.direct = 2;
    this.y += this.speed;
}

Tank.prototype.moveLeft = function () {
    this.direct = 3;
    this.x -= this.speed;
}

// Wall define wall
function Wall(x, y) {
    this.x = x;
    this.y = y;
    //   this.wall_number=wall_number;
    this.isLive = true;
}
// module1 -- draw walls 
var module1 = {
    a: function () {
        var wall_num = [];
        return wall_num;
    },
    // lateral
    draw_x: function (x, y, wall_number) {
        var wall_1 = [];
        for (var i = 0; i < wall_number; i++) {
            wall_1[i] = new Wall(x + 40 * i, y);
        }
        return wall_1;
    },
    // vertical
    draw_y: function (x, y, wall_number) {
        var wall_2 = new Array();
        for (var i = 0; i < wall_number; i++) {
            wall_2[i] = new Wall(x, y + 40 * i);
        }
        return wall_2;
    },
    init: function (wall_num) {
        wall_num.push(module1.draw_x(150, 280, 5));//成组构造

        wall_num.push(module1.draw_x(110, 400, 7));
        wall_num.push(module1.draw_y(70, 400, 5));
        wall_num.push(module1.draw_y(390, 400, 5));

        wall_num.push(module1.draw_x(190, 480, 3));
        wall_num.push(module1.draw_y(150, 480, 4));
        wall_num.push(module1.draw_y(310, 480, 4));


        return wall_num;
    },
    draw_wall: function (wall) {
        var drawing = document.getElementById("tankMap");
        var cxt = drawing.getContext('2d');
        var drawing2 = document.getElementById("heroMap2");
        var cxt2 = drawing2.getContext('2d');
        var img = new Image();
        img.src = "./img/wall_1.png";
        function draw() {
            var walls = []; var wall_num = wall;
            for (var i = 0; i < wall_num.length; i++) {
                for (var j = 0; j < wall_num[i].length; j++) {
                    if (wall_num[i][j].isLive == true) {
                        cxt.drawImage(img, wall_num[i][j].x, wall_num[i][j].y);
                        walls.push(wall_num[i][j])
                    }
                }
            }
            return (walls);
        }
        // onload = anonymous function
        img.onload = function () {
            draw();
        }
        return draw;
    }
}

// Bullet define bullet and run
function Bullet(x, y, direct, speed, type, tank) {
    this.x = x;
    this.y = y;
    this.isLive = true;
    this.timer = timer;
    this.direct = direct;
    this.speed = speed;
    this.type = type;
    this.tank = tank;
}
Bullet.prototype.run = function () {
    if (this.x < 0 || this.x > 500 || this.y < 10 || this.y > 600 || this.isLive == false) {
        clearInterval(this.timer);
        this.isLive = false;
        if (this.type == "enemy") {
            this.tank.bulletIsLive = false;
        }
    } else {
        switch (this.direct) {
            case 0:
                this.y -= this.speed;
                break;
            case 1:
                this.x += this.speed;
                break;
            case 2:
                this.y += this.speed;
                break;
            case 3:
                this.x -= this.speed;
                break;
        }
    }
}


// module2 -- draw tank
var drawing = document.getElementById("tankMap");
var cxt = drawing.getContext('2d');
var module2 = (function () {
    var drawTank = function (tank) {
        if (tank.isLive == true) {
            switch (tank.direct) {
                case 0:
                case 2:
                    if (tank.type == "hero") {
                        cxt.fillStyle = "white";
                    } else if (tank.type == "enemy") {
                        cxt.fillStyle = "green";
                    }
                    // tank body
                    cxt.fillRect(tank.x, tank.y, 5, 30);
                    cxt.fillRect(tank.x + 15, tank.y, 5, 30);
                    cxt.fillRect(tank.x + 6, tank.y + 5, 8, 20);

                    // tank head
                    if (tank.type == "hero") {
                        cxt.fillStyle = "red";
                    } else if (tank.type == "enemy") {
                        cxt.fillStyle = "red";
                    }
                    cxt.arc(tank.x + 10, tank.y + 15, 4, 0, 2 * Math.PI, true);
                    cxt.fill();
                    if (tank.type == "hero") {
                        cxt.strokeStyle = "red";
                    } else if (tank.type == "enemy") {
                        cxt.strokeStyle = "red";
                    }
                    cxt.lineWidth = 2;
                    cxt.beginPath();
                    cxt.moveTo(tank.x + 10, tank.y + 15);
                    if (tank.direct == 0) {
                        cxt.lineTo(tank.x + 10, tank.y);
                    } else if (tank.direct == 2) {
                        cxt.lineTo(tank.x + 10, tank.y + 30)
                    }
                    cxt.closePath();
                    cxt.stroke();
                    break;
                case 1:
                case 3:
                    if (tank.type == "hero") {
                        cxt.fillStyle = "white";
                    } else if (tank.type == "enemy") {
                        cxt.fillStyle = "green";
                    }
                    cxt.fillRect(tank.x, tank.y, 30, 5);
                    cxt.fillRect(tank.x, tank.y + 15, 30, 5);
                    cxt.fillRect(tank.x + 5, tank.y + 6, 20, 8);
                    if (tank.type == "hero") {
                        cxt.fillStyle = "red";
                    } else if (tank.type == "enemy") {
                        cxt.fillStyle = "red";
                    }
                    cxt.arc(tank.x + 15, tank.y + 10, 4, 0, 2 * Math.PI, true);
                    cxt.fill();
                    if (tank.type == "hero") {
                        cxt.strokeStyle = "red";
                    } else if (tank.type == "enemy") {
                        cxt.strokeStyle = "red";
                    }
                    cxt.lineWidth = 2;
                    cxt.beginPath();
                    cxt.moveTo(tank.x + 15, tank.y + 10);
                    if (tank.direct == 1) {
                        cxt.lineTo(tank.x + 30, tank.y + 10);
                    } else if (tank.direct == 3) {
                        cxt.lineTo(tank.x, tank.y + 10)
                    }
                    cxt.closePath();
                    cxt.stroke();
                    break;
            }
        }
    }
    return {
        m1: drawTank
    }
})();


// Hero define hero tank
function Hero(x, y, direct, speed, type) {
    Tank.call(this, x, y, direct, speed, type);
}
Hero.prototype = new Tank();
Hero.prototype.constructor = Hero;

// hero bullet shot
Hero.prototype.shotEnermy = function () {
    switch (this.direct) {
        case 0:
            heroBullet = new Bullet(this.x + 9, this.y, this.direct, 10);
            break;
        case 1:
            heroBullet = new Bullet(this.x + 30, this.y + 9, this.direct, 10);
            break;
        case 2:
            heroBullet = new Bullet(this.x + 9, this.y + 30, this.direct, 10);
            break;
        case 3:
            heroBullet = new Bullet(this.x, this.y + 9, this.direct, 10);
            break;
    }

    if (heroBullets.length < 18) {
        heroBullets.push(heroBullet);
    }
    console.log(heroBullets.length);
    heroBullet.timer = timer

    // hero bullet run
    var timer = window.setInterval("heroBullets[" + (heroBullets.length - 1) + "].run()", 50);
}


// EnemyTank define enemy tank
function EnemyTank(x, y, direct, speed, type) {
    Tank.call(this, x, y, direct, speed, type);
    this.isLive = true;
    this.count = 0;
    this.bulletIsLive = true;
}
EnemyTank.prototype = new Tank();
EnemyTank.prototype.constructor = EnemyTank;
EnemyTank.prototype.run = function (wall) {

    // flag indicates the positional relationship between tank and wall 
    // random generate tank new direction
    var flag = null;
    switch (this.direct) {
        case 0:
            for (var i = 0; i < wall.length; i++) {
                if ((this.y - 30 > 0) && ((this.y - this.speed) >= wall[i].y) && ((this.y - this.speed) < wall[i] + 40) && (this.x >= wall[i].x) && (this.x < wall[i].x + 40) && (wall[i].isLive == true)) {
                    flag = 1;
                }
            }

            if ((this.y < 125) && (this.x < 80) || (this.y < 325) && (this.y > 280) && (this.x < 470) && (this.x > 350)) {
                flag = 1;
            }

            if (this.y > 0 && flag != 1) {
                this.y -= this.speed;
            } else {
                this.direct = Math.round(Math.random() * 3);
            }
            break;
        case 1:
            for (var i = 0; i < wall.length; i++) {
                if ((this.x + 20 < 500) && (this.y + 30 > wall[i].y) && (this.y < (wall[i].y + 40)) && ((this.x + 30 + this.speed) >= wall[i].x) && ((this.x + this.speed) < (wall[i].x + 40)) && (wall[i].isLive == true)) {
                    flag = 1;
                }
            }
            if ((this.x > 315) && (this.x < 400) && (this.y > 260) && (this.y < 320)) {
                flag = 1;
            }


            if (this.x + 30 < 500 && flag != 1) {
                this.x += this.speed;
            } else {
                this.direct = Math.round(Math.random() * 3);
            }
            break;
        case 2:
            for (var i = 0; i < wall.length; i++) {
                if ((this.y + 30 < 600) && ((this.y + 30 + this.speed) >= wall[i].y) && (this.y + 30 + this.speed) < (wall[i].y + 40) && (this.x + 20 >= wall[i].x) && (this.x < wall[i].x + 40) && (wall[i].isLive == true)) {
                    flag = 1;
                }
            }
            if ((this.x > 330) && (this.x < 470) && (this.y > 245) && (this.y < 320)) {
                flag = 1;
            }

            if (this.y + 30 < 600 && flag != 1) {
                this.y += this.speed;
            } else {
                this.direct = Math.round(Math.random() * 3);
            }
            break;
        case 3:
            for (var i = 0; i < wall.length; i++) {
                if ((this.x - 20 > 0) && (this.y + 30 > wall[i].y) && (this.y < (wall[i].y + 40)) && ((this.x - this.speed) >= wall[i].x) && ((this.x - this.speed) < (wall[i].x + 40)) && (wall[i].isLive == true)) {
                    flag = 1;
                }
            }
            if ((this.y < 125) && (this.x < 85) || (this.y < 320) && (this.y > 255) && (this.x < 475) && (this.x > 350)) {
                flag = 1;
            }

            if (this.x > 0 && flag != 1) {
                this.x -= this.speed;
            } else {
                this.direct = Math.round(Math.random() * 3);
            }
            break;
    }
    if (this.count > 30) {
        this.direct = Math.round(Math.random() * 3);
        this.count = 0;
    }
    this.count++;
}


var $ = function (ele) {
    return document.querySelector(ele);
}

var result = module1.init(module1.a());
module1.draw_wall(result);

//var result2=module11.init(module11.a());
//module11.draw_water(result2);

// var result3=module10.init(module10.a());
// module10.draw_grass(result2);

var heroBullets = new Array();
var hero = new Hero(230, 340, 0, 5, "hero");
var enemyBullets = new Array();
var enemyTanks = new Array();

// module3 -- draw tank and start move by flash
var module3 = (function (draw, enemy_number, walls, hero) {
    var drawHero = function () {
        draw(hero);
    }
    var Enermy = function () {
        for (var i = 0; i < 4; i++) {
            var enemy_x = (Math.random() * 490);
            var enemy_y = Math.random() * 350;
            var enemy_direct = Math.round(Math.random() * 3);
            var enemyTank = new EnemyTank(enemy_x, enemy_y, enemy_direct, 2, "enemy");
            enemyTanks[i] = enemyTank;
            // enemy tank run
            window.setInterval((function (j, wallstmp) {
                return function () {
                    enemyTanks[j].run(wallstmp);
                }
            })(i, walls), 50);
            // bullet
            if (enemyTanks[i].isLive == true) {
                switch (enemyTanks[i].direct) {
                    case 0:
                        var eb = new Bullet(enemyTanks[i].x + 9, enemyTanks[i].y, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                        break;
                    case 1:
                        var eb = new Bullet(enemyTanks[i].x + 30, enemyTanks[i].y + 9, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                        break;
                    case 2:
                        var eb = new Bullet(enemyTanks[i].x + 9, enemyTanks[i].y + 30, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                        break;
                    case 3: //右
                        var eb = new Bullet(enemyTanks[i].x, enemyTanks[i].y + 9, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                        break;
                }
            }
            enemyBullets[i] = eb;
            // enemy bullet run
            var ettimer = window.setInterval("enemyBullets[" + i + "].run()", 50);
            enemyBullets[i].timer = ettimer;
        }
    }
    var drawEnermy = function () {
        for (var i = 0; i < 4; i++) {
            draw(enemyTanks[i]);
        }
    }
    return {
        m1: drawHero,
        m2: Enermy,
        m3: drawEnermy
    }
}
)(module2.m1, 4, module1.draw_wall(result)(), hero);
module3.m2();

// module4 -- hero tank move and shot
var module4 = (function (wall, hero) {
    var test = function (evt) {
        evt = (evt) ? evt : ((window.event) ? window.event : "");
        var code = evt.keyCode ? evt.which : evt.keyCode;
        switch (code) {
            case 87:
                // W move up  hero.y = hero.y - hero.speed
                for (var i = 0; i < wall.length; i++) {
                    if ((hero.y - 30 > 0) && ((hero.y - hero.speed) >= wall[i].y) && (hero.y - hero.speed) < (wall[i].y + 40) && (hero.x >= wall[i].x) && (hero.x < wall[i].x + 40) && (wall[i].isLive == true)) {
                        flag = 1;
                    }
                }

                if ((hero.y < 125) && (hero.x < 80) || (hero.y < 325) && (hero.y > 280) && (hero.x < 470) && (hero.x > 350)) {
                    flag = 1;
                }

                if ((hero.y - hero.speed) > 0 && flag != 1) {
                    hero.moveUp();
                }
                break;

            case 68:
                // D move right hero.x = hero.x + hero.speed;
                for (var i = 0; i < wall.length; i++) {
                    if ((hero.x + 20 < 500) && (hero.y + 30 > wall[i].y) && (hero.y < (wall[i].y + 40)) && ((hero.x + 30 + hero.speed) >= wall[i].x) && ((hero.x + hero.speed) < (wall[i].x + 40)) && (wall[i].isLive == true)) {
                        flag = 1;
                    }
                }
                if ((hero.x > 315) && (hero.x < 400) && (hero.y > 260) && (hero.y < 320)) {
                    flag = 1;
                }

                if ((hero.x + hero.speed) < 500 && flag != 1) {
                    hero.moveRight();
                }
                break;

            case 83:
                // S move down hero.y = hero.y + hero.speed
                for (var i = 0; i < wall.length; i++) {
                    if ((hero.y + 30 < 600) && ((hero.y + 30 + hero.speed) >= wall[i].y) && (hero.y + 30 + hero.speed) < (wall[i].y + 40) && (hero.x + 20 >= wall[i].x) && (hero.x < wall[i].x + 40) && (wall[i].isLive == true)) {
                        flag = 1;
                    }
                }

                if ((hero.x > 330) && (hero.x < 470) && (hero.y > 245) && (hero.y < 320)) {
                    flag = 1;
                }

                if ((hero.y + hero.speed) < 600 && flag != 1) {
                    hero.moveDown();
                }
                break;

            case 65:
                // A move left hero.x = hero.x - hero.speed;
                var flag = null;
                for (var i = 0; i < wall.length; i++) {
                    if ((hero.x - 20 > 0) && (hero.y + 30 > wall[i].y) && (hero.y < (wall[i].y + 40)) && ((hero.x - hero.speed) >= wall[i].x) && ((hero.x - hero.speed) < (wall[i].x + 40)) && (wall[i].isLive == true)) {
                        flag = 1;
                    }
                }

                if ((hero.y < 125) && (hero.x < 85) || (hero.y < 320) && (hero.y > 255) && (hero.x < 475) && (hero.x > 350)) {
                    flag = 1;
                }

                if ((hero.x - hero.speed) > 0 && flag != 1) {
                    hero.moveLeft();
                }
                break;

            case 74:
                // J shot
                hero.shotEnermy();
                break;
        }
        module9.m1();
    }
    return {
        m1: test
    }
})(module1.draw_wall(result)(), hero)


// module5 -- draw hero bullet, enemy bullet and enemyBullet run
var drawing = document.getElementById("heroMap2");
var cxt2 = drawing.getContext("2d");
var module5 = (function () {
    var drawHeroBullet = function () {
        for (var i = 0; i < heroBullets.length; i++) {
            var heroBullet = heroBullets[i];
            if (heroBullet != null && heroBullet.isLive) {
                cxt2.fillStyle = "yellow";
                cxt2.fillRect(heroBullet.x, heroBullet.y, 2, 2);
            }
        }
    }
    var drawEnemyBullet = function (tank, i, bullet) {
        var etBullet = bullet;
        if (tank.bulletIsLive == true && tank.isLive == true) {
            cxt.fillStyle = "#00FEFE";
            cxt.fillRect(etBullet.x, etBullet.y, 2, 2);
        }
    }
    var EnemyBullet_logic = function () {
        for (var i = 0; i < enemyTanks.length; i++) {
            if (enemyTanks[i].bulletIsLive == false) {
                enemyTanks[i].bulletIsLive = true;
                if (enemyTanks[i].isLive == true) {
                    switch (enemyTanks[i].direct) {
                        case 0:
                            enemyBullets[i] = new Bullet(enemyTanks[i].x + 9, enemyTanks[i].y, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                            break;
                        case 1:
                            enemyBullets[i] = new Bullet(enemyTanks[i].x + 30, enemyTanks[i].y + 9, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                            break;
                        case 2:
                            enemyBullets[i] = new Bullet(enemyTanks[i].x + 9, enemyTanks[i].y + 30, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                            break;
                        case 3: //右
                            enemyBullets[i] = new Bullet(enemyTanks[i].x, enemyTanks[i].y + 9, enemyTanks[i].direct, 2, "enemy", enemyTanks[i]);
                            break;
                    }
                }
                var ettimer = window.setInterval("enemyBullets[" + i + "].run()", 50);
                enemyBullets[i].timer = ettimer;
                (function (enemyTanks, j, enemyBullets) {
                    return module5.m2(enemyTanks, i, enemyBullets);
                })(enemyTanks[i], i, enemyBullets[i]);

            }
            else {
                (function (enemyTanks, j, enemyBullets) {
                    return module5.m2(enemyTanks, i, enemyBullets);
                })(enemyTanks[i], i, enemyBullets[i]);
            }
        }
    }
    return {
        m1: drawHeroBullet,
        m2: drawEnemyBullet,
        m3: EnemyBullet_logic
    }
}
)()

// module6 -- flash bullet
var module6 = (function (drawHeroBullet) {
    var flashBullet = function () {
        cxt2.globalAlpha = 1;
        cxt2.clearRect(0, 0, 500, 600);
        drawHeroBullet();
    }
    return {
        m1: flashBullet
    }
}
)(module5.m1)

// God define Boss
function God(x, y) {
    this.x = x;
    this.y = y;
    this.isLive = true;
}

//Bomb define bomb 
function Bomb(x, y) {
    this.x = x;
    this.y = y;
    this.isLive = true;
    
    this.blood = 9;

    this.bloodDown = function () {
        if (this.blood > 0) {
            this.blood--;
        } else {
            this.isLive = false;
        }
    }
}

// module7 -- bullet hit tank event
var bombs = new Array();
var god = new God(230, 564);
var module7 = {
    drawGod: function () {
        var img = new Image()
        img.src = "./img/boss.png";
        cxt.drawImage(img, god.x, god.y);
    },
    drawEnemyBomb: function () {
        for (var i = 0; i < bombs.length; i++) {
            //取出一颗炸弹
            var bomb = bombs[i];
            if (bomb.isLive) {
                //更据当前这个炸弹的生命值，来画出不同的炸弹图片
                if (bomb.blood > 6) {  //显示最大炸弹图
                    var img1 = new Image();
                    img1.src = "./img/bomb_1.gif";
                    var x = bomb.x;
                    var y = bomb.y;
                    img1.onload = function () {
                        cxt.drawImage(img1, x, y, 30, 30);
                    }
                } else if (bomb.blood > 3) {
                    var img2 = new Image();
                    img2.src = "./img/bomb_2.gif";
                    var x = bomb.x;
                    var y = bomb.y;
                    img2.onload = function () {
                        cxt.drawImage(img2, x, y, 30, 30);
                    }
                } else {
                    var img3 = new Image();
                    img3.src = "./img/bomb_3.gif";
                    var x = bomb.x;
                    var y = bomb.y;
                    img3.onload = function () {
                        cxt.drawImage(img3, x, y, 30, 30);
                    }
                }
                //减血
                bomb.bloodDown();
                if (bomb.blood <= 0) {
                    //怎么办?把这个炸弹从数组中去掉
                    bombs.splice(i, 1);
                }
            }
        }
    },
    isHitEnemyTank: function (wall) {
        //取出每颗子弹
        for (var i = 0; i < heroBullets.length; i++) {
            //取出一颗子弹
            var heroBullet = heroBullets[i];
            /* alert(heroBullet.islive);*/
            if (heroBullet.isLive === true) { //子弹是活的，才去判断
                //让这颗子弹去和遍历每个墙判断
                for (var i = 0; i < wall.length; i++) {
                    var walls = wall[i];
                    if (walls.isLive == true) {
                        if (heroBullet.x >= wall[i].x && heroBullet.x <= wall[i].x + 40
                            && heroBullet.y >= wall[i].y && heroBullet.y <= wall[i].y + 40) {
                            wall[i].isLive = false; console.log(wall[i].x, wall[i].isLive);
                            heroBullet.isLive = false;
                        }
                    }
                }
                //让这颗子弹去和遍历每个敌人坦克判断
                for (var j = 0; j < enemyTanks.length; j++) {
                    var enemyTank = enemyTanks[j];
                    if (enemyTank.isLive === true) {
                        //子弹击中敌人坦克的条件是什么? 很多思路 , 韩老师的思想是
                        //(看看这颗子弹，是否进入坦克所在矩形)
                        //根据当时敌人坦克的方向来决定
                        switch (enemyTank.direct) {
                            case 0: //敌人坦克向上
                            case 2://敌人坦克向下
                                if (heroBullet.x >= enemyTank.x && heroBullet.x <= enemyTank.x + 20
                                    && heroBullet.y >= enemyTank.y && heroBullet.y <= enemyTank.y + 30) {
                                    //把坦克isLive 设为false ,表示死亡
                                    enemyTank.isLive = false;
                                    //该子弹也死亡
                                    heroBullet.isLive = false;
                                    //创建一颗炸弹
                                    var bomb = new Bomb(enemyTank.x, enemyTank.y);
                                    //然后把该炸弹放入到bombs数组中
                                    bombs.push(bomb);
                                }
                                break;
                            case 1: //敌人坦克向右
                            case 3://敌人坦克向左
                                if (heroBullet.x >= enemyTank.x && heroBullet.x <= enemyTank.x + 30
                                    && heroBullet.y >= enemyTank.y && heroBullet.y <= enemyTank.y + 20) {
                                    //把坦克isLive 设为false ,表示死亡
                                    enemyTank.isLive = false;
                                    heroBullet.isLive = false;
                                    //创建一颗炸弹
                                    var bomb = new Bomb(enemyTank.x, enemyTank.y);
                                    //然后把该炸弹放入到bombs数组中
                                    bombs.push(bomb);
                                }
                                break;
                        }
                    }
                }
            }
        }
    },//判断敌人子弹是否击中英雄坦克
    isHitHeroTank: function (wall) {
        //取出每颗子弹
        for (var i = 0; i < enemyTanks.length; i++) {
            var etBullet = enemyBullets[i];
            //这里，我们加入了一句话，但是要知道这里加，是需要对整个程序有把握
            if (etBullet.isLive == true) {
                //打到老大
                if (god.isLive == true) {
                    if (etBullet.x >= god.x && etBullet.x <= god.x + 40
                        && etBullet.y >= god.y && etBullet.y <= god.y + 40) {
                        god.isLive = false;
                        etBullet.isLive = false;
                        alert("You have lost");
                        clearInterval(timer);
                    }
                }
                //让这颗子弹去和遍历每个墙判断
                for (var j = 0; j < wall.length; j++) {//j不能换成i
                    var walls = wall[j];
                    if (walls.isLive == true) {
                        if (etBullet.x >= wall[j].x && etBullet.x <= wall[j].x + 40
                            && etBullet.y >= wall[j].y && etBullet.y <= wall[j].y + 40) {
                            wall[j].isLive = false;
                            etBullet.isLive = false;
                        }
                    }
                }
                if (hero.isLive == true) {
                    switch (hero.direct) {
                        case 0: //敌人坦克向上
                        case 2://敌人坦克向下
                            if (etBullet.x >= hero.x && etBullet.x <= hero.x + 20
                                && etBullet.y >= hero.y && etBullet.y <= hero.y + 30) {
                                //把坦克isLive 设为false ,表示死亡
                                hero.isLive = false;
                                //该子弹也死亡
                                etBullet.isLive = false;
                                //创建一颗炸弹
                                var bomb = new Bomb(hero.x, hero.y);
                                //然后把该炸弹放入到bombs数组中
                                bombs.push(bomb);
                                alert("You have lost");
                                clearInterval(timer);
                            }
                            break;
                        case 1: //敌人坦克向右
                        case 3://敌人坦克向左
                            if (etBullet.x >= hero.x && etBullet.x <= hero.x + 30
                                && etBullet.y >= hero.y && etBullet.y <= hero.y + 20) {
                                //把坦克isLive 设为false ,表示死亡
                                hero.isLive = false;
                                etBullet.isLive = false;
                                alert("You have lost");
                                clearInterval(timer);
                                //创建一颗炸弹
                                var bomb = new Bomb(hero.x, hero.y);
                                //然后把该炸弹放入到bombs数组中
                                bombs.push(bomb);
                            }
                            break;
                    }
                }
            }
        }
    }//画出敌人炸弹
}
// module8 -- victory decision
var module8 = {
    Success: function () {
        if (enemyTanks[0].isLive == false && enemyTanks[1].isLive == false && enemyTanks[2].isLive == false && enemyTanks[3].isLive == false) {
            document.getElementById("heroMap3").style.display = "block";
        }
    }
}


//module9 -- main()
var module9 = (function () {
    var flash = function () {
        cxt.clearRect(0, 0, 500, 600);
        module3.m1();// hero
        module3.m3();// enemy

        var result2 = module11.init(module11.a());
        module11.draw_water(result2)(); //water

        var result3 = module10.init(module10.a());
        module10.draw_grass(result3); //grass

        module1.draw_wall(result)();//wall
        
        module7.drawGod();//god

        module5.m3();//enemy bullet
        module7.isHitEnemyTank(module1.draw_wall(result)());//hero hit enemy
        module7.isHitHeroTank(module1.draw_wall(result)());//enemy hit hero
        module7.drawEnemyBomb();//boom
        module8.Success();
    }
    return {
        m1: flash
    }
})()


// module 10,11 redraw grass and water
function Grass(x, y) {
    this.x = x;
    this.y = y;
    this.isLive = true;
}

var module10 = {
    a: function () {
        var grass_num = [];
        return grass_num;
    },
    draw_x: function (x, y, grass_number) {
        var grass_1 = [];
        for (var i = 0; i < grass_number; i++) {
            grass_1[i] = new Grass(x + 40 * i, y);
        }
        return grass_1;
    },
    draw_y: function (x, y, grass_number) {
        var grass_2 = new Array();
        for (var i = 0; i < grass_number; i++) {
            grass_2[i] = new Grass(x, y + 40 * i);
        }
        return grass_2;
    },
    init: function (grass_number) {
        grass_number.push(module10.draw_x(30, 280, 3));

        grass_number.push(module10.draw_y(420, 0, 3));
        grass_number.push(module10.draw_y(460, 0, 3));

        return grass_number;
    },
    draw_grass: function (grass) {
        var drawing = document.getElementById("tankMap");
        var cxt = drawing.getContext('2d');
        var drawing2 = document.getElementById("heroMap2");
        var cxt2 = drawing2.getContext('2d');
        var img = new Image();
        img.src = "./img/grass.png";
        function draw() {
            var grasss = []; var grass_number = grass;
            for (var i = 0; i < grass_number.length; i++) {
                for (var j = 0; j < grass_number[i].length; j++) {
                    if (grass_number[i][j].isLive == true) {
                        cxt.drawImage(img, grass_number[i][j].x, grass_number[i][j].y);
                        grasss.push(grass_number[i][j])
                    }
                }
            }
            return (grasss);
        }
        img.onload = function () {
            draw();
        }
        return draw;
    }
}

function Water(x, y) {
    this.x = x;
    this.y = y;
    this.isLive = true;
}

var module11 = {
    a: function () {
        var water_num = [];
        return water_num;
    },
    draw_x: function (x, y, water_number) {
        var water_1 = [];
        for (var i = 0; i < water_number; i++) {
            water_1[i] = new Water(x + 40 * i, y);
        }
        return water_1;
    },
    draw_y: function (x, y, water_number) {
        var water_2 = new Array();
        for (var i = 0; i < water_number; i++) {
            water_2[i] = new Water(x, y + 40 * i);
        }
        return water_2;
    },
    init: function (water_number) {
        water_number.push(module11.draw_x(350, 280, 3));

        water_number.push(module11.draw_y(0, 0, 3));
        water_number.push(module11.draw_y(40, 0, 3));

        return water_number;
    },
    draw_water: function (water) {
        var drawing = document.getElementById("tankMap");
        var cxt = drawing.getContext('2d');
        var drawing2 = document.getElementById("heroMap2");
        var cxt2 = drawing2.getContext('2d');
        var img = new Image();
        img.src = "./img/water.png";
        function draw() {
            var waters = []; var water_number = water;
            for (var i = 0; i < water_number.length; i++) {
                for (var j = 0; j < water_number[i].length; j++) {
                    if (water_number[i][j].isLive == true) {
                        cxt.drawImage(img, water_number[i][j].x, water_number[i][j].y);
                        waters.push(water_number[i][j])
                    }
                }
            }
            return (waters);
        }
        img.onload = function () {
            draw();
        }
        return draw;
    }
}

var timer = setInterval("module9.m1()", 100);
var timer2 = setInterval("module6.m1()", 100);