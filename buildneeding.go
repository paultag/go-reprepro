/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@debian.org>, 2015
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

package reprepro

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type BuildNeedingPackage struct {
	Source   string
	Version  string
	Location string
	Arch     string
}

func (repo *Repo) BuildNeeding(suite string, arch string, glob *string) ([]BuildNeedingPackage, error) {
	ret := []BuildNeedingPackage{}
	args := []string{"build-needing", suite, arch}
	if glob != nil {
		args = append(args, *glob)
	}
	cmd := repo.Command(args...)
	out, err := cmd.Output()
	if err != nil {
		return ret, err
	}
	reader := bufio.NewReader(strings.NewReader(string(out)))
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, " \n\r\t")
		if err != nil {
			if err == io.EOF {
				break
			}
			return ret, err
		}
		els := strings.Split(line, " ")
		if len(els) == 4 {
			ret = append(ret, BuildNeedingPackage{
				Source:   els[0],
				Version:  els[1],
				Location: els[2],
				Arch:     els[3],
			})
		} else if len(els) == 3 {
			ret = append(ret, BuildNeedingPackage{
				Source:   els[0],
				Version:  els[1],
				Location: els[2],
				Arch:     arch,
			})
		} else {
			return nil, fmt.Errorf("Unknown line: %s\n", line)
		}
	}
	return ret, nil
}

// vim: foldmethod=marker
