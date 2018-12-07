<?php

/**
 * Created by PhpStorm.
 * @Author     : likingfit@likingfit.com
 * @CreateTime 2018/12/7 16:39:17
 */
class MyArray {
    private $data;
    private $capacity;
    private $length;

    public function __construct($capacity) {
        if ($capacity <= 0) {
            return null;
        }
        $this->data = [];
        $this->capacity = $capacity;
        $this->length = 0;
    }

    //判断是否为空
    public function isEmpty() {
        if ($this->length == 0) {
            return true;
        }

        return false;
    }

    //判断是否已满
    public function checkIsFull() {
        if ($this->length == $this->capacity) {
            return true;
        }

        return false;
    }

    //判断是否越界
    public function checkOutOfRange($index) {
        if ($index > (($this->capacity) - 1) && $index < 0) {
            return true;
        } else {
            return false;
        }
    }

    //插入
    /**
     * 最好：O（1）
     * 最坏：O（n）
     * 平均：O（n）
     */
    public function insert($index, $data) {
        if ($this->checkIsFull()) {
            return 1;
        }

        if ($this->checkOutOfRange($index)) {
            return 2;
        }
//复杂度：O（n）  ，如果是无序的，可以把$index的值搬到最后一个。这样复杂度变为O（1）

        var_dump($index.'---'.$data);


        for ($i = $index; $i < $this->length ; $i++) {
            $this->data[$i+1] = $this->data[$i];
        }

        $this->data[$index] = $data;
        $this->length++;

        return 0;
    }

    //删除
    public function delete($index) {
        if ($this->checkOutOfRange($index)) {
            return 2;
        }

        for ($i = $index; $i < ($this->length - 1); $i++) {
            $this->data[$i] = $this->data[$i + 1];
        }


        $this->length--;

        return 0;
    }

    //find
    public function find($index) {
        if ($this->checkOutOfRange($index)) {
            return [
                2,
                0
            ];
        }
        echo $this->data[$index];

        return $this->data[$index];
    }

    //print
    public function printf() {
        $format = " ";
        for ($i = 0; $i < $this->length; $i++) {
            $dat = $this->data[$i];
            $format .= "|" . $dat;
        }
        echo "\n" . $this->length . "-------" . $format . "\n";
    }
}


$myArr = new MyArray(10);

$myArr->insert(0, 1);
$myArr->insert(1, 2);
$myArr->insert(2, 2);
$myArr->insert(3, 6);
$myArr->insert(4, 5);
$myArr->insert(5, 4);
$myArr->insert(6, 3);

$myArr->printf();

$myArr->find(3);
$myArr->delete(3);

$myArr->printf();

$myArr->insert(7, 8);
$myArr->printf();